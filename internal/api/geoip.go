// SPDX-License-Identifier: GPL-3.0-or-later

package api

import (
	"net/http"

	"github.com/fjcrux/mieru-web-ui/internal/geoip"
)

type geoipResponse struct {
	Datasets       []geoip.Dataset      `json:"datasets"`
	Categories     []geoip.Category     `json:"categories"`
	SiteDatasets   []geoip.Dataset      `json:"siteDatasets"`
	SiteCategories []geoip.SiteCategory `json:"siteCategories"`
}

func (s *Server) geoipState() geoipResponse {
	datasets, _ := s.Geo.Datasets()
	categories, _ := s.Geo.Categories()
	siteDatasets, _ := s.Geo.SiteDatasets()
	siteCategories, _ := s.Geo.SiteCategories()
	if datasets == nil {
		datasets = []geoip.Dataset{}
	}
	if categories == nil {
		categories = []geoip.Category{}
	}
	if siteDatasets == nil {
		siteDatasets = []geoip.Dataset{}
	}
	if siteCategories == nil {
		siteCategories = []geoip.SiteCategory{}
	}
	return geoipResponse{
		Datasets: datasets, Categories: categories,
		SiteDatasets: siteDatasets, SiteCategories: siteCategories,
	}
}

func (s *Server) handleGeoipList(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.geoipState())
}

type addDatasetRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	// Kind selects the dataset type: "geosite" for domain lists, anything
	// else (including empty) means geoip.
	Kind string `json:"kind"`
}

func (s *Server) handleGeoipAddDataset(w http.ResponseWriter, r *http.Request) {
	var req addDatasetRequest
	if !readJSON(w, r, &req) {
		return
	}
	var err error
	if req.Kind == "geosite" {
		err = s.Geo.AddSiteDataset(r.Context(), req.Name, req.URL)
	} else {
		err = s.Geo.AddDataset(r.Context(), req.Name, req.URL)
	}
	if err != nil {
		writeErr(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, s.geoipState())
}

func (s *Server) handleGeoipDeleteDataset(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.URL.Query().Get("kind") == "geosite" {
		err = s.Geo.DeleteSiteDataset(r.PathValue("name"))
	} else {
		err = s.Geo.DeleteDataset(r.PathValue("name"))
	}
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, s.geoipState())
}
