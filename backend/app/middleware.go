package app

import (
	"net/http"
)

var APIUsage uint64

func EnforceMethod(method string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// current := atomic.AddUint64(&APIUsage, 1)
		// fmt.Printf("%d API used: %s %s\n", current, r.Method, r.URL.Path)

		if r.Method != method {
			respondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		h(w, r)
	}
}
