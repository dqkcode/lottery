package lottery

import (
	"encoding/json"
	"lottery/crawler"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type HandlerFunc struct {
}

//create new HandlerFunc
func NewHandlerFunc() *HandlerFunc {
	return &HandlerFunc{}
}

func (h *HandlerFunc) getResult(w http.ResponseWriter, r *http.Request) {

	numbers := strings.Split(r.FormValue("numbers"), "")
	chanel := r.FormValue("chanel")
	time := r.FormValue("time")
	crawler.GetResultFromWebsite(numbers, chanel, time)

	err := json.NewEncoder(w).Encode([]map[string]interface{}{
		{
			"numbers": numbers,
			"chanel":  chanel,
			"time":    time,
		},
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
func (h *HandlerFunc) getTest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	w.Write([]byte("id send: " + id))
}
