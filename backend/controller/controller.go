package controller

import (
	"encoding/json"
	"net/http"

	"web-imagecomparison/model"
	"web-imagecomparison/service"
)

type ProjectController struct {
	Service *service.ProjectService
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewProjectController(service *service.ProjectService) *ProjectController {
	return &ProjectController{Service: service}
}

func (vc *ProjectController) HandleEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, ErrorResponse{"method not allowed"})
		return
	}

	var v model.ProjectModel
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		respondJSON(w, http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if err := vc.Service.CreateEntry(&v); err != nil {
		respondJSON(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
		return
	}

	respondJSON(w, http.StatusCreated, v)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
