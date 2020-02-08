package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/method"
	"trainpix-api/api/method/train"
)

func Route(port int, logger *logrus.Logger) {
	logger.Trace("Creating router")
	router := mux.NewRouter()

	methods := []method.Method{
		train.Get{Group: "train", Method: "get"},
	}

	for _, currentMethod := range methods {
		router.HandleFunc("/v1/" + currentMethod.GetGroup() + "/" + currentMethod.GetMethod(), func(writer http.ResponseWriter, request *http.Request) {
			HandleAPI(writer, request, currentMethod, logger)
		})
	}

	logger.Debug("Creating HTTP server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		logger.Panic("Server creation error: ", err)
	}
}

func HandleAPI(writer http.ResponseWriter, request *http.Request, method method.Method, logger *logrus.Logger) {
	writer.Header().Add("content-type", "application/json")

	query := request.URL.Query()
	encoder := json.NewEncoder(writer)

	logger.Debug("API Request: /" + method.GetGroup() + "/" + method.GetMethod())

	response, err := method.Process(query)

	if err == nil {
		writer.WriteHeader(200)
	} else if err.Error() == "400" {
		writer.WriteHeader(400)
	} else if err.Error() == "404" {
		writer.WriteHeader(404)
	} else if err.Error() == "500" {
		writer.WriteHeader(500)
	}

	encoder.Encode(response)
}
