package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/method"
	"trainpix-api/api/method/depot"
	"trainpix-api/api/method/photo"
	"trainpix-api/api/method/train"
)

func Route(port int, logger *logrus.Logger) {
	logger.Trace("Creating router")
	router := mux.NewRouter()

	methods := []method.Method{
		train.Get{Group: "train", Method: "get"},
		train.Search{Group: "train", Method: "search"},
		train.QSearch{Group: "train", Method: "qsearch"},

		photo.Get{Group: "photo", Method: "get"},
		photo.Random{Group: "photo", Method: "random"},
		photo.Search{Group: "photo", Method: "search"},

		depot.Get{Group: "depot", Method: "get"},
		depot.Search{Group: "depot", Method: "search"},
	}

	router.HandleFunc("/v1/{group}/{method}", func(writer http.ResponseWriter, request *http.Request) {
		HandleAPI(writer, request, methods, logger)
	})

	logger.Debug("Creating HTTP server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		logger.Panic("Server creation error: ", err)
	}
}

func HandleAPI(writer http.ResponseWriter, request *http.Request, methods []method.Method, logger *logrus.Logger) {
	vars := mux.Vars(request)
	query := request.URL.Query()
	encoder := json.NewEncoder(writer)

	writer.Header().Add("content-type", "application/json")

	for _, current := range methods {
		if current.GetGroup() == vars["group"] && current.GetMethod() == vars["method"] {
			logger.Debug("API Request: /" + current.GetGroup() + "/" + current.GetMethod() + " " + query.Encode())
			result := current.Process(query)
			writer.WriteHeader(result.GetStatus())
			encoder.Encode(result)
			return
		}
	}

	fmt.Fprint(writer, "{}")
}
