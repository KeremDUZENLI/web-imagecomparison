package app

import "net/http"

func NewRouter(svc *ProjectService) http.Handler {
	mux := http.NewServeMux()
	ctrl := NewProjectController(svc)
	mux.Handle("/", http.FileServer(http.Dir("../docs")))
	mux.HandleFunc("/api/votes", EnforceMethod(http.MethodPost, ctrl.HandleVotes))
	mux.HandleFunc("/api/ratings", EnforceMethod(http.MethodGet, ctrl.HandleRatings))
	return mux
}
