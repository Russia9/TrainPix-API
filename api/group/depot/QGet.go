package depot

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/parser"
)

func QGet(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()
	id := 1

	if v.Get("id") != "" {
		id, _ = strconv.Atoi(v.Get("id"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		id = -1
	}

	count := 5
	if v.Get("count") != "" {
		count, _ = strconv.Atoi(v.Get("count"))
	}

	quick := true
	if v.Get("quick") == "0" {
		quick = false
	}

	if !quick && count > 20 {
		count = 20
	}

	logger.Debug("depot/get: id='", id, "'")

	depot, trains, err := parser.DepotGet(id, count, quick)
	if err != nil {
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Trace(err)
		}
	}

	json.NewEncoder(w).Encode(response.DepotGet{
		Depot:  depot,
		Trains: trains,
	})
}
