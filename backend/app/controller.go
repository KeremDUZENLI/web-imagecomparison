package app

import (
	"encoding/json"
	"net/http"
)

type ProjectController struct {
	Service *ProjectService
}

func NewProjectController(svc *ProjectService) *ProjectController {
	return &ProjectController{Service: svc}
}

func (pc *ProjectController) HandleVotes(w http.ResponseWriter, r *http.Request) {
	var dto VoteDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
		return
	}

	vote, err := pc.Service.ProcessVote(r.Context(), &dto)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not process vote"})
		return
	}

	ratings, err := pc.Service.GetAllRatings(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not fetch ratings"})
		return
	}

	respondJSON(w, http.StatusCreated, map[string]any{
		"vote":    vote,
		"ratings": ratings,
	})
}

func (pc *ProjectController) HandleRatings(w http.ResponseWriter, r *http.Request) {
	ratings, err := pc.Service.GetAllRatings(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to get all ratings"})
		return
	}
	respondJSON(w, http.StatusOK, ratings)
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
