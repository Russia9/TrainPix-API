package photo

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"trainpix-api/api/response"
	"trainpix-api/parser"
)

func Random(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
	w.Header().Add("content-type", "application/json")

	logger.Debug("photo/random")

	photo, train, err := parser.RandomPhotoGet()
	if err != nil {
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Trace(err)
		}
	}

	json.NewEncoder(w).Encode(response.PhotoRandomGet{
		Photo: photo,
		Train: train,
	})
}
