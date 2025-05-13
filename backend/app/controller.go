package app

import (
	"encoding/json"
	"net/http"
)

type ProjectController struct {
	Service *ProjectService
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewProjectController(service *ProjectService) *ProjectController {
	return &ProjectController{Service: service}
}

func (vc *ProjectController) HandleVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var v ProjectModel
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if err := vc.Service.RecordVote(&v); err != nil {
		respondError(w, http.StatusInternalServerError, "failed to record vote")
		return
	}

	if err := vc.Service.RecalculateRatings(&v); err != nil {
		respondError(w, http.StatusInternalServerError, "failed to update ratings")
		return
	}

	respondJSON(w, http.StatusCreated, v)
}

func (vc *ProjectController) HandleRatings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	m, err := vc.Service.GetAllRatings()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "failed to get all ratings")
		return
	}
	respondJSON(w, http.StatusOK, m)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}
