package train

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/parser"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()
	query := "ЭР2"
	if v.Get("query") != "" {
		query = v.Get("query")
	}
	count := 5
	if v.Get("count") != "" {
		count, _ = strconv.Atoi(v.Get("count"))
	}

	resultCode := 200
	trains, err := parser.TrainSearch(query, count, false)
	if err != nil {
		if err.Error() == "404" {
			resultCode = 200
		} else {
			resultCode = 500
		}
	}

	json.NewEncoder(w).Encode(response.TrainSearch{
		ResultCode: resultCode,
		Trains:     trains,
	})
}

func QuickSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()
	query := "ЭР2"
	if v.Get("query") != "" {
		query = v.Get("query")
	}
	count := 5
	if v.Get("count") != "" {
		count, _ = strconv.Atoi(v.Get("count"))
	}
	trains, _ := parser.TrainSearch(query, count, true)
	json.NewEncoder(w).Encode(trains)
}
