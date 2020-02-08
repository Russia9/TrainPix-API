package photo

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/photo"
	"trainpix-api/parse"
)

type Get struct {
	Group  string
	Method string
}

func (object Get) GetGroup() string {
	return object.Group
}

func (object Get) GetMethod() string {
	return object.Method
}

func (object Get) Process(params url.Values) response.Response {
	id := -1
	var err error

	if params.Get("id") != "" {
		id, err = strconv.Atoi(params.Get("id"))

		if err != nil {
			return photo.Get{Status: 400}
		}
	} else {
		return photo.Get{Status: 400}
	}

	result, err := parse.PhotoGet(id, false)

	if err != nil {
		if err.Error() == "404" {
			return photo.Get{Status: 404}
		} else {
			return photo.Get{Status: 500}
		}
	}

	return photo.Get{Status: 200, Result: result}
}
