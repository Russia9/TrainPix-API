package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/group/photo"
	"trainpix-api/api/group/train"
)

func Api(port int, logger *logrus.Logger) {
	logger.Trace("Creating router")
	router := mux.NewRouter()

	// Train API Group
	router.HandleFunc("/api/v0.4/train/get", func(w http.ResponseWriter, r *http.Request) {
		train.Get(w,r, logger)
	})

	router.HandleFunc("/api/v0.4/train/search", func(w http.ResponseWriter, r *http.Request) {
		train.Search(w,r, logger)
	})
	router.HandleFunc("/api/v0.4/train/qsearch", func(w http.ResponseWriter, r *http.Request) {
		train.QuickSearch(w,r, logger)
	})

	// Photo API Group
	router.HandleFunc("/api/v0.4/photo/get", func(w http.ResponseWriter, r *http.Request) {
		photo.Get(w, r, logger)
	})
	router.HandleFunc("/api/v0.4/photo/random", func(w http.ResponseWriter, r *http.Request) {
		photo.Random(w, r, logger)
	})


	logger.Debug("Creating HTTP server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		logger.Panic("Server creation error: ", err)
	}
}
