// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	pb "github.com/enfein/mieru/v3/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"

	"github.com/fjcrux/mieru-web-ui/internal/sharelink"
)

// --- consumer side: chain this panel's egress through upstream peers ---

func (s *Server) handleChainList(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.Peers.List())
}

type addPeerRequest struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func (s *Server) handleChainAdd(w http.ResponseWriter, r *http.Request) {
	var req addPeerRequest
	if !readJSON(w, r, &req) {
		return
	}
	if err := s.Peers.Add(req.Name, req.Key); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	// The peer is now available as an egress proxy; push it to mita.
	if err := s.applyEgress(r.Context()); err != nil {
		writeErr(w, http.StatusBadGateway, "peer added but applying egress failed: "+err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, s.Peers.List())
}

func (s *Server) handleChainRemove(w http.ResponseWriter, r *http.Request) {
	if err := s.Peers.Remove(r.PathValue("name")); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := s.applyEgress(r.Context()); err != nil {
		writeErr(w, http.StatusBadGateway, "peer removed but applying egress failed: "+err.Error())
		return
	}
	writeJSON(w, http.StatusOK, s.Peers.List())
}

// --- provider side: hand out a peer key so another panel can chain to us ---

type chainKeyResponse struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// handleChainKey creates a dedicated user on this server and returns a
// mieru:// key another panel imports to route its egress through us.
func (s *Server) handleChainKey(w http.ResponseWriter, r *http.Request) {
	host, err := s.Store.Setting("public_host")
	if err != nil || host == "" {
		writeErr(w, http.StatusConflict, "public host is not configured - set it in Settings")
		return
	}

	name := "peer-" + randHex(4)
	password := randHex(16)

	_, err = s.Mita.UpdateConfig(r.Context(), func(cfg *pb.ServerConfig) error {
		cfg.Users = append(cfg.Users, &pb.User{
			Name:     proto.String(name),
			Password: proto.String(password),
		})
		s.applyPortInvariant(cfg)
		return nil
	})
	if err != nil {
		writeMitaErr(w, err)
		return
	}
	if err := s.Store.SetUserSecret(name, password); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	s.reconcileProxy(r.Context(), false)

	cfg, err := s.Mita.GetConfig(r.Context())
	if err != nil {
		writeMitaErr(w, err)
		return
	}
	links, err := sharelink.Build(sharelink.Params{
		Username:     name,
		Password:     password,
		Host:         host,
		PortBindings: cfg.GetPortBindings(),
		MTU:          cfg.GetMtu(),
	})
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, chainKeyResponse{Name: name, Key: links.MieruURL})
}

func randHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
