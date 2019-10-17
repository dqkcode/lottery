package types

import "net/http"

type (
	Route struct {
		Path        string
		Method      string
		HandlerFunc http.HandlerFunc
		Queries     []string
		Middlewares []Middleware
	}
	Middleware func(h http.Handler) http.Handler
)
