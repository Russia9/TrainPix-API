package train

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

	if v.Get("id") != "" {
		id, _ = strconv.Atoi(v.Get("id"))
	}

	logger.Debug("train/get: id='", id, "'")

	train, err := parser.TrainGet(id, false)
	resultCode := 200
	if err != nil {
		if err.Error() == "404" {
			resultCode = 404
		} else {
			resultCode = 500
		}
	}

	json.NewEncoder(w).Encode(response.TrainGet{
		ResultCode: resultCode,
		Train:      train,
	})
}
