// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"
)

var userNameRe = regexp.MustCompile(`^[a-zA-Z0-9._-]{1,64}$`)

// onlineWindow is how recently a user must have transferred data to count as
// "online". mita has no per-session user attribution, so activity is inferred
// from the per-user traffic counters' time-series history.
const onlineWindow = 5 * time.Minute

// lastActiveMillis returns the most recent Unix-millis timestamp at which the
// user transferred any data, derived from the history of their traffic
// counters (upload/download bytes are COUNTER_TIME_SERIES). Returns 0 when the
// user has never had traffic.
func lastActiveMillis(item *pb.UserWithMetrics) int64 {
	var latest int64
	for _, m := range item.GetMetrics() {
		for _, h := range m.GetHistory() {
			// Only windows that actually recorded traffic mark activity.
			if h.GetDelta() > 0 && h.GetTimeUnixMilli() > latest {
				latest = h.GetTimeUnixMilli()
			}
		}
	}
	return latest
}

// Quota mirrors mita's per-user quota (traffic cap over a rolling window).
type Quota struct {
	Days      int32 `json:"days"`
	Megabytes int32 `json:"megabytes"`
}

type userInfo struct {
	Name           string  `json:"name"`
	Quotas         []Quota `json:"quotas"`
	AllowPrivateIP bool    `json:"allowPrivateIP"`
	HasSecret      bool    `json:"hasSecret"`
	// HasSubscription is true when a permanent subscription link is issued.
	HasSubscription bool `json:"hasSubscription"`
	// Traffic metrics from GetUsers; zero when the user has no traffic.
	Metrics map[string]int64 `json:"metrics"`
	// LastActiveUnixMs is when the user last transferred data (0 = never),
	// inferred from traffic-counter history since mita can't map sessions to
	// users. The UI shows "online" when this is within the last few minutes.
	LastActiveUnixMs int64 `json:"lastActiveUnixMs"`
}

func (s *Server) handleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.Mita.Users(r.Context())
	if err != nil {
		if proxyNotReady(err) {
			writeJSON(w, http.StatusOK, []userInfo{})
			return
		}
		writeMitaErr(w, err)
		return
	}
	secrets, err := s.Store.UserSecretNames()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	subs, err := s.Store.SubTokenUsernames()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	out := []userInfo{}
	for _, item := range users.GetItems() {
		u := item.GetUser()
		info := userInfo{
			Name:            u.GetName(),
			Quotas:          []Quota{},
			AllowPrivateIP:  u.GetAllowPrivateIP(),
			HasSecret:       secrets[u.GetName()],
			HasSubscription: subs[u.GetName()],
			Metrics:         map[string]int64{},
		}
		for _, q := range u.GetQuotas() {
			info.Quotas = append(info.Quotas, Quota{Days: q.GetDays(), Megabytes: q.GetMegabytes()})
		}
		for _, m := range item.GetMetrics() {
			info.Metrics[m.GetName()] = m.GetValue()
		}
		info.LastActiveUnixMs = lastActiveMillis(item)
		out = append(out, info)
	}
	writeJSON(w, http.StatusOK, out)
}

type upsertUserRequest struct {
	Password       string  `json:"password"`
	Quotas         []Quota `json:"quotas"`
	AllowPrivateIP bool    `json:"allowPrivateIP"`
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		upsertUserRequest
	}
	if !readJSON(w, r, &req) {
		return
	}
	if !userNameRe.MatchString(req.Name) {
		writeErr(w, http.StatusBadRequest, "invalid user name (allowed: letters, digits, . _ -, max 64)")
		return
	}
	if len(req.Password) < 8 {
		writeErr(w, http.StatusBadRequest, "password must be at least 8 characters")
		return
	}
	_, err := s.Mita.UpdateConfig(r.Context(), func(cfg *pb.ServerConfig) error {
		for _, u := range cfg.GetUsers() {
			if u.GetName() == req.Name {
				return fmt.Errorf("user %q already exists", req.Name)
			}
		}
		cfg.Users = append(cfg.Users, buildUser(req.Name, req.upsertUserRequest))
		// Ports must accompany users in the same write (mita rejects users
		// without ports).
		s.applyPortInvariant(cfg)
		return nil
	})
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := s.Store.SetUserSecret(req.Name, req.Password); err != nil {
		writeErr(w, http.StatusInternalServerError, "user created in mita but saving secret failed: "+err.Error())
		return
	}
	// Start proxying if this is the first user (and ports exist).
	s.reconcileProxy(r.Context(), false)
	writeJSON(w, http.StatusCreated, map[string]bool{"ok": true})
}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	var req upsertUserRequest
	if !readJSON(w, r, &req) {
		return
	}
	if req.Password != "" && len(req.Password) < 8 {
		writeErr(w, http.StatusBadRequest, "password must be at least 8 characters")
		return
	}
	_, err := s.Mita.UpdateConfig(r.Context(), func(cfg *pb.ServerConfig) error {
		for i, u := range cfg.GetUsers() {
			if u.GetName() != name {
				continue
			}
			updated := buildUser(name, req)
			if req.Password == "" {
				// Keep the existing hash; mita preserves it on round-trip.
				updated.Password = nil
				updated.HashedPassword = proto.String(u.GetHashedPassword())
			}
			cfg.Users[i] = updated
			return nil
		}
		return fmt.Errorf("user %q not found", name)
	})
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.Password != "" {
		if err := s.Store.SetUserSecret(name, req.Password); err != nil {
			writeErr(w, http.StatusInternalServerError, "password updated in mita but saving secret failed: "+err.Error())
			return
		}
		// Old share links carried the previous password; invalidate them.
		// Subscription tokens survive: they resolve the current password at
		// fetch time, so clients pick up the change on their next refresh.
		_ = s.Store.DeleteShareTokensForUser(name)
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	_, err := s.Mita.UpdateConfig(r.Context(), func(cfg *pb.ServerConfig) error {
		users := cfg.GetUsers()
		for i, u := range users {
			if u.GetName() == name {
				cfg.Users = append(users[:i], users[i+1:]...)
				// Drop ports too if that was the last user.
				s.applyPortInvariant(cfg)
				return nil
			}
		}
		return fmt.Errorf("user %q not found", name)
	})
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	_ = s.Store.DeleteUserSecret(name)
	_ = s.Store.DeleteShareTokensForUser(name)
	_ = s.Store.DeleteSubTokenForUser(name)
	// Stop proxying if that was the last user, to avoid a mita crashloop.
	s.reconcileProxy(r.Context(), false)
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func buildUser(name string, req upsertUserRequest) *pb.User {
	u := &pb.User{
		Name:           proto.String(name),
		AllowPrivateIP: proto.Bool(req.AllowPrivateIP),
	}
	if req.Password != "" {
		u.Password = proto.String(req.Password)
	}
	for _, q := range req.Quotas {
		u.Quotas = append(u.Quotas, &pb.Quota{
			Days:      proto.Int32(q.Days),
			Megabytes: proto.Int32(q.Megabytes),
		})
	}
	return u
}
