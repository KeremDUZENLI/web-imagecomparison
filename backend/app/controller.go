package app

import (
	"encoding/json"
	"net/http"
)

type ProjectController struct {
	service ProjectService
}

func NewProjectController(svc ProjectService) *ProjectController {
	return &ProjectController{service: svc}
}

func (pc *ProjectController) HandleUsers(w http.ResponseWriter, r *http.Request) {
	users, err := pc.service.GetAllUserNames(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not fetch users"})
		return
	}
	respondJSON(w, http.StatusOK, users)
}

func (pc *ProjectController) HandleVotes(w http.ResponseWriter, r *http.Request) {
	var votesDto VotesDTO
	if err := json.NewDecoder(r.Body).Decode(&votesDto); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
		return
	}

	votes, err := pc.service.PostVote(r.Context(), &votesDto)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not process vote"})
		return
	}

	ratings, err := pc.service.GetAllRatings(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not fetch ratings"})
		return
	}

	respondJSON(w, http.StatusCreated, map[string]any{
		"votes":   votes,
		"ratings": ratings,
	})
}

func (pc *ProjectController) HandleRatings(w http.ResponseWriter, r *http.Request) {
	ratings, err := pc.service.GetAllRatings(r.Context())
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
