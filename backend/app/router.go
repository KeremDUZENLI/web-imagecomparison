package app

import "net/http"

func NewRouter(svc *ProjectService) http.Handler {
	controller := NewProjectController(svc)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../docs")))
	mux.HandleFunc("/api/votes", EnforceMethod(http.MethodPost, controller.HandleVotes))
	mux.HandleFunc("/api/ratings", EnforceMethod(http.MethodGet, controller.HandleRatings))

	return mux
}
