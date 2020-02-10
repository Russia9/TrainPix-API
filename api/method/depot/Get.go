package depot

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/depot"
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
	var result *object.Depot
	trainCount := 5
	photoCount := 5
	quick := false

	if params.Get("id") == "" {
		return depot.Get{Status: 400}
	} else {
		id, err = strconv.Atoi(params.Get("id"))
		if err != nil {
			return depot.Get{Status: 400}
		}
	}

	if params.Get("trainCount") != "" {
		trainCount, err = strconv.Atoi(params.Get("trainCount"))
		if err != nil {
			return depot.Get{Status: 400}
		}
	}

	if params.Get("photoCount") != "" {
		photoCount, err = strconv.Atoi(params.Get("photoCount"))
		if err != nil {
			return depot.Get{Status: 400}
		}
	}

	if params.Get("quick") != "" {
		quickInt, err := strconv.Atoi(params.Get("quick"))
		if err != nil {
			return depot.Get{Status: 400}
		}
		if quickInt == 1 {
			quick = true
		}
	}

	result, err = parse.DepotGet(id, trainCount, photoCount, quick)
	if err != nil {
		if err.Error() == "404" {
			return depot.Get{Status: 404}
		} else {
			return depot.Get{Status: 500}
		}
	}

	return depot.Get{
		Status: 200,
		Result: result,
	}
}
