package app

import (
	"net/http"
)

func NewRouter(pc *ProjectController, cfg MiddlewareConfig) http.Handler {
	apiMux := http.NewServeMux()
	registerAPIRoutes(apiMux, pc)

	logHandler := LogRequest(cfg)(apiMux)
	cacheHandler := CacheControl(cfg)(http.FileServer(http.Dir("../")))

	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", logHandler))
	router.Handle("/", cacheHandler)

	return router
}

func registerAPIRoutes(mux *http.ServeMux, pc *ProjectController) {
	mux.HandleFunc("/users", EnforceMethod(http.MethodGet, pc.HandleUsers))
	mux.HandleFunc("/votes", EnforceMethod(http.MethodPost, pc.HandleVotes))
	mux.HandleFunc("/ratings", EnforceMethod(http.MethodGet, pc.HandleRatings))
}
