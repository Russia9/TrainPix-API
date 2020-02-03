package photo

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/parser"
)

func Get(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()
	id := 1
	quick := false

	if v.Get("id") != "" {
		id, _ = strconv.Atoi(v.Get("id"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		id=0
	}

	if v.Get("quick") != "" {
		if v.Get("quick") == "1" {
			quick = true
		}
	}

	logger.Debug("photo/get: id='", id, "' quick='", quick, "'")

	photo, err := parser.PhotoGet(id, quick)
	if err != nil {
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Trace(err)
		}
	}

	json.NewEncoder(w).Encode(response.PhotoGet{
		Photo: photo,
	})
}
