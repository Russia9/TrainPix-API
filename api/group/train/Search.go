package train

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/parser"
)

func Search(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
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

	logger.Debug("train/search: query='", query, "' count='", count, "'")

	resultCode := 200
	trains, err := parser.TrainSearch(query, count, false)
	if err != nil {
		if err.Error() == "404" {
			resultCode = 200
		} else {
			resultCode = 500
		}
		logger.Trace(err)
	}

	json.NewEncoder(w).Encode(response.TrainSearch{
		ResultCode: resultCode,
		Trains:     trains,
	})
}

func QuickSearch(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
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
	logger.Debug("train/qsearch: query='", query, "' count='", count, "'")

	trains, err := parser.TrainSearch(query, count, true)
	resultCode := 200
	if err != nil {
		if err.Error() == "404" {
			resultCode = 404
		} else {
			resultCode = 500
		}
		logger.Trace(err)
	}

	json.NewEncoder(w).Encode(response.TrainSearch{
		ResultCode: resultCode,
		Trains:     trains,
	})
}
