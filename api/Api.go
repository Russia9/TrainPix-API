package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"trainpix-api/api/group/photo"
	"trainpix-api/api/group/train"
)

func Api(port int, log *logrus.Logger) {
	log.Trace("Creating router")
	router := mux.NewRouter()

	// Train API Group
	router.HandleFunc("/api/v0.1/train/get", func(w http.ResponseWriter, r *http.Request) {
		train.Get(w,r,log)
	})

	router.HandleFunc("/api/v0.1/train/search", func(w http.ResponseWriter, r *http.Request) {
		train.Search(w,r,log)
	})
	router.HandleFunc("/api/v0.1/train/qsearch", func(w http.ResponseWriter, r *http.Request) {
		train.QuickSearch(w,r,log)
	})

	// Photo API Group
	router.HandleFunc("/api/v0.1/photo/get", func(w http.ResponseWriter, r *http.Request) {
		photo.Get(w, r, log)
	})

	log.Debug("Creating HTTP server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		log.Panic("Server creation error: ", err)
	}
}
