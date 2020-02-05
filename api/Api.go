package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	//query := r.URL.Query()

	logger.Debug("API Request: /" + vars["group"] + "/" + vars["method"])

	switch vars["group"] {
	case "train":
		switch vars["method"] {
		case "search":
			break
		case "get":
			break
		}
		break
	case "photo":
		switch vars["method"] {
		case "search":
			break
		case "get":
			break
		case "random":
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
}
