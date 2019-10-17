package lottery

import (
	"lottery/types"
	"net/http"
)

func (h *HandlerFunc) Routes() []types.Route {
	return []types.Route{
		{
			Path:        "/result",
			Method:      http.MethodGet,
			HandlerFunc: h.getResult,
		},
		{
			Path:        "/result",
			Method:      http.MethodPost,
			HandlerFunc: h.getResult,
		},
		{
			Path:        "/result/{numbers:[\\d]}/{chanel:[\\w]}/{time:[\\w]}",
			Method:      http.MethodGet,
			HandlerFunc: h.getResult,
		},
		// tinh=Hà nội&ngay=1&thang=3&nam=2017
		// {
		// 	Path:        "/result/{numbers:[a-zA-Z0-9]+}/{chanel:[a-zA-Z]+/{day:[0-9]+/{mon:[0-9]+/{year:[0-9]+}",
		// 	Method:      http.MethodGet,
		// 	HandlerFunc: h.getResult,
		// },
	}
}
