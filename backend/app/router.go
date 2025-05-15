package app

import (
	"net/http"
)

func NewRouter(pc *ProjectController, cfg MiddlewareConfig) http.Handler {
	apiMux := http.NewServeMux()
	registerAPIRoutes(apiMux, pc)

	handler := LogRequest(cfg)(apiMux)
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", handler))

	router.Handle("/", http.FileServer(http.Dir("../")))
	return router
}

func registerAPIRoutes(mux *http.ServeMux, pc *ProjectController) {
	mux.HandleFunc("/users", EnforceMethod(http.MethodGet, pc.HandleUsers))
	mux.HandleFunc("/votes", EnforceMethod(http.MethodPost, pc.HandleVotes))
	mux.HandleFunc("/ratings", EnforceMethod(http.MethodGet, pc.HandleRatings))
}
