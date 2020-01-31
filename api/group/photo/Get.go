package photo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/parser"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()
	id := 1
	quick := false

	if v.Get("id") != "" {
		id, _ = strconv.Atoi(v.Get("id"))
	}

	if v.Get("quick") != "" {
		if v.Get("quick") == "1" {
			quick = true
		}
	}

	photo, err := parser.PhotoGet(id, quick)
	resultCode := 200
	if err != nil {
		if err.Error() == "404" {
			resultCode = 404
		} else {
			resultCode = 500
		}
	}

	json.NewEncoder(w).Encode(response.PhotoGet{
		ResultCode: resultCode,
		Photo:      photo,
	})
}