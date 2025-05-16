package app

import (
	"fmt"
	"net/http"
	"path"
	"strings"
	"sync/atomic"
	"time"
)

type MiddlewareConfig struct {
	EnableLogging      bool
	DisableStaticCache bool
}

var apiCounter uint64

func EnforceMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			respondJSON(w, http.StatusMethodNotAllowed,
				map[string]string{"error": "method not allowed"})
			return
		}
		handler(w, r)
	}
}

func LogRequest(cfg MiddlewareConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if cfg.EnableLogging {
				count := atomic.AddUint64(&apiCounter, 1)
				ts := time.Now().Format("2006/01/02 15:04:05")
				fmt.Printf("%d API used | %s | %s %s\n", count, ts, r.Method, r.URL.Path)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func CacheControl(cfg MiddlewareConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if cfg.DisableStaticCache {
				ext := strings.ToLower(path.Ext(r.URL.Path))
				switch ext {
				case ".js", ".css", ".html":
					w.Header().Set("Cache-Control", "no-store, must-revalidate")
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
