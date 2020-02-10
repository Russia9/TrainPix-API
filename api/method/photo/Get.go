package photo

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/photo"
	"trainpix-api/object"
	"trainpix-api/parse"
)

type Get struct {
	Group  string
	Method string
}

func (method Get) GetGroup() string {
	return method.Group
}

func (method Get) GetMethod() string {
	return method.Method
}

func (method Get) Process(params url.Values) response.Response {
	var err error
	var id int
	var result *object.Photo
	quick := false

	if params.Get("id") == "" {
		return photo.Get{Status: 400}
	} else {
		id, err = strconv.Atoi(params.Get("id"))
		if err != nil {
			return photo.Get{Status: 400}
		}
	}

	if params.Get("quick") != "" {
		quickInt, err := strconv.Atoi(params.Get("quick"))
		if err != nil {
			return photo.Get{Status: 400}
		}
		if quickInt == 1 {
			quick = true
		}
	}

	result, err = parse.PhotoGet(id, quick)
	if err != nil {
		if err.Error() == "404" {
			return photo.Get{Status: 404}
		} else {
			return photo.Get{Status: 500}
		}
	}

	return photo.Get{
		Status: 200,
		Result: result,
	}
}
