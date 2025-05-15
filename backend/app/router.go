package app

import (
	"log"
	"net/http"
)

func NewRouter(pc *ProjectController) http.Handler {
	apiMux := http.NewServeMux()
	registerAPIRoutes(apiMux, pc)

	rootMux := http.NewServeMux()
	rootMux.Handle("/api/", http.StripPrefix("/api", LoggingMiddleware(apiMux)))

	fs := http.FileServer(http.Dir("../"))
	rootMux.Handle("/", fs)

	return rootMux
}

func registerAPIRoutes(mux *http.ServeMux, pc *ProjectController) {
	mux.HandleFunc("/users", EnforceMethod(http.MethodGet, pc.HandleUsers))
	mux.HandleFunc("/votes", EnforceMethod(http.MethodPost, pc.HandleVotes))
	mux.HandleFunc("/ratings", EnforceMethod(http.MethodGet, pc.HandleRatings))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
