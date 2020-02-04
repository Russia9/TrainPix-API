package railway

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Search(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()

	query := "Московская железная дорога"
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

	logger.Debug("railway/search: query='", query, "' count='", count, "'")

	/*trains, countFound, countParsed, err := parser.TrainSearch(query, count, false, nil)
	if err != nil {
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Trace(err)
		}
	}*/

	json.NewEncoder(w).Encode(nil)
}
