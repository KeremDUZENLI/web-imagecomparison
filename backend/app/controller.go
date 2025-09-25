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

func (pc *ProjectController) HandleSurvey(w http.ResponseWriter, r *http.Request) {
	var surveyModel SurveysModel
	if err := json.NewDecoder(r.Body).Decode(&surveyModel); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid payload")
		return
	}
	if err := pc.service.PostSurvey(r.Context(), surveyModel); err != nil {
		errorJSON(w, http.StatusInternalServerError, "could not post survey")
		return
	}
	respondJSON(w, http.StatusCreated, surveyModel)
}

func (pc *ProjectController) HandleUsers(w http.ResponseWriter, r *http.Request) {
	users, err := pc.service.GetAllUsernames(r.Context())
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "could not get all usernames")
		return
	}
	if users == nil {
		users = []string{}
	}
	respondJSON(w, http.StatusOK, users)
}

func (pc *ProjectController) HandleVotes(w http.ResponseWriter, r *http.Request) {
	var votesDto VotesDTO
	if err := json.NewDecoder(r.Body).Decode(&votesDto); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid payload")
		return
	}

	votes, err := pc.service.PostVote(r.Context(), &votesDto)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "could not post vote")
		return
	}

	ratings, err := pc.service.GetAllRatings(r.Context())
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, "could not get all ratings")
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
		errorJSON(w, http.StatusBadRequest, "could not get all ratings")
		return
	}
	respondJSON(w, http.StatusOK, ratings)
}

func errorJSON(w http.ResponseWriter, status int, msg string) {
	respondJSON(w, status, map[string]string{"error": msg})
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
