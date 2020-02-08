package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Route(port int, logger *logrus.Logger) {
	logger.Trace("Creating router")
	router := mux.NewRouter()

	router.HandleFunc("/v1/{group}/{method}", func(writer http.ResponseWriter, request *http.Request) {
		HandleAPI(writer, request, logger)
	})

	logger.Debug("Creating HTTP server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		logger.Panic("Server creation error: ", err)
	}
}

func HandleAPI(writer http.ResponseWriter, request *http.Request, logger *logrus.Logger)  {
	writer.Header().Add("content-type", "application/json")

	vars := mux.Vars(request)
	query := request.URL.Query()
	encoder := json.NewEncoder(writer)

	logger.Debug("API Request: /" + vars["group"] + "/" + vars["method"])
}