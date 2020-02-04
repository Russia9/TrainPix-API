package railway

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()
	id := 1

	if v.Get("id") != "" {
		id, _ = strconv.Atoi(v.Get("id"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		id = -1
	}

	logger.Debug("railway/get: id='", id, "'")

	/*train, err := parser.TrainGet(id, false)
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
