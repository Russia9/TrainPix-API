package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/groups/photo"
	"trainpix-api/api/groups/train"
)

func Api(port int, logger *logrus.Logger) {
	logger.Trace("Creating router")
	router := mux.NewRouter()

	router.HandleFunc("/v1/{group}/{method}", func(w http.ResponseWriter, r *http.Request) {
		HandleAPI(w, r, logger)
	})

	logger.Debug("Creating HTTP server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		logger.Panic("Server creation error: ", err)
	}
}

func HandleAPI(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
	w.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	query := r.URL.Query()
	encoder := json.NewEncoder(w)
	var err error

	logger.Debug("API Request: /" + vars["group"] + "/" + vars["method"])

	switch vars["group"] {
	case "train":
		switch vars["method"] {
		case "search":
			_, err = train.Search(query, false)
			break
		case "qsearch":
			_, err = train.Search(query, true)
			break
		case "get":
			_, err = train.Get(query)
			break
		}
		break
	case "photo":
		switch vars["method"] {
		case "search":
			break
		case "get":
			_, err = photo.Get(query)
			break
		case "random":
			_, err = photo.RandomGet()
			break
		}
		break
	case "railway":
		switch vars["method"] {
		case "search":
			break
		case "get":
			break
		}
		break
	case "depot":
		switch vars["method"] {
		case "search":
			break
		case "get":
			break
		}
		break
	default:
		break
	}

	if err != nil {
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == "400" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	err = encoder.Encode(nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
