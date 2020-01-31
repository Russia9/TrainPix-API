package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"trainpix-api/api/group/photo"
	"trainpix-api/api/group/train"
)

func Api(port int) {
	router := mux.NewRouter()

	// Train API Group
	router.HandleFunc("/api/v0.1/train/get", train.Get)

	router.HandleFunc("/api/v0.1/train/search", train.Search)
	router.HandleFunc("/api/v0.1/train/qsearch", train.QuickSearch)

	// Photo API Group
	router.HandleFunc("/api/v0.1/photo/get", photo.Get)

	router.HandleFunc("/api/v0.1/photo/search", photo.Search)

	_ = http.ListenAndServe(":"+strconv.Itoa(port), router)
}
