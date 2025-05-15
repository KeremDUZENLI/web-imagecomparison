package app

import "net/http"

func NewRouter(pc *ProjectController) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../")))
	mux.HandleFunc("/api/users", EnforceMethod(http.MethodGet, pc.HandleUsers))
	mux.HandleFunc("/api/votes", EnforceMethod(http.MethodPost, pc.HandleVotes))
	mux.HandleFunc("/api/ratings", EnforceMethod(http.MethodGet, pc.HandleRatings))

	return mux
}
