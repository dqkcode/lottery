package server

import (
	"log"
	"net/http"
)

// LoggerMiddleware log the resquest, response
func LoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("path: %s, method: %s", r.URL.Path, r.Method)
			h.ServeHTTP(w, r)
		})
}
