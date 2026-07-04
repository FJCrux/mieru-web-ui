// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"net/http"

	"github.com/fjcrux/mieru-web-ui/internal/geoip"
)

type geoipResponse struct {
	Datasets   []geoip.Dataset  `json:"datasets"`
	Categories []geoip.Category `json:"categories"`
}

func (s *Server) geoipState() geoipResponse {
	datasets, _ := s.Geo.Datasets()
	categories, _ := s.Geo.Categories()
	if datasets == nil {
		datasets = []geoip.Dataset{}
	}
	if categories == nil {
		categories = []geoip.Category{}
	}
	return geoipResponse{Datasets: datasets, Categories: categories}
}

func (s *Server) handleGeoipList(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.geoipState())
}

type addDatasetRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (s *Server) handleGeoipAddDataset(w http.ResponseWriter, r *http.Request) {
	var req addDatasetRequest
	if !readJSON(w, r, &req) {
		return
	}
	if err := s.Geo.AddDataset(r.Context(), req.Name, req.URL); err != nil {
		writeErr(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, s.geoipState())
}

func (s *Server) handleGeoipDeleteDataset(w http.ResponseWriter, r *http.Request) {
	if err := s.Geo.DeleteDataset(r.PathValue("name")); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, s.geoipState())
}
