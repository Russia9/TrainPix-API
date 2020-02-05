package train

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
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

	if count > 20 {
		count = 20
	}

	logger.Debug("train/search: query='", query, "' count='", count, "'")

	trains, countFound, countParsed, err := parser.TrainSearch(query, count, false, getParams(v))
	if err != nil {
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Trace(err)
		}
	}

	json.NewEncoder(w).Encode(response.TrainSearch{
		Found:  countFound,
		Parsed: countParsed,
		TrainList: trains,
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

	trains, countFound, countParsed, err := parser.TrainSearch(query, count, true, getParams(v))
	if err != nil {
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Trace(err)
		}
	}

	json.NewEncoder(w).Encode(response.TrainSearch{
		Found:  countFound,
		Parsed: countParsed,
		TrainList: trains,
	})
}

func getParams(v url.Values) map[string]string {
	params := make(map[string]string)

	if v.Get("state") != "" {
		params["state"] = v.Get("state")
	}
	if v.Get("order") != "" {
		params["order"] = v.Get("order")
	}
	if v.Get("st") != "" {
		params["st"] = v.Get("st")
	}
	return params
}
