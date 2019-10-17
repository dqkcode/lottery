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
	}
}
