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

	// numbers := strings.Split(r.FormValue("numbers"), "")
	// chanel := r.FormValue("chanel")
	// time := r.FormValue("time")
	query := r.URL.Query()
	numbers := strings.Split(query["numbers"][0], "")
	channel := query["channel"][0]
	day := query["day"][0]
	month := query["month"][0]
	year := query["year"][0]

	crawler.GetResultFromWebsite(numbers, channel, day, month, year)

	err := json.NewEncoder(w).Encode([]map[string]interface{}{
		{
			"numbers": numbers,
			"channel": channel,
			"day":     day,
			"month":   month,
			"year":    year,
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
