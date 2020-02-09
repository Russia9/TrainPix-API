package depot

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/depot"
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
			return depot.Get{Status: 400}
		}
	} else {
		return depot.Get{Status: 400}
	}

	count := 5
	if params.Get("count") != "" {
		count, err = strconv.Atoi(params.Get("count"))
		if err != nil {
			return depot.Get{Status: 400}
		}
	}

	if count > 20 {
		count = 20
	}

	result, err := parse.DepotGet(id, count, false)

	if err != nil {
		if err.Error() == "404" {
			return depot.Get{Status: 404}
		}
		return depot.Get{Status: 500}
	}

	return depot.Get{Status: 200, Result: result}
}
