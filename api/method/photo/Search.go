package photo

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/photo"
	"trainpix-api/parse"
)

type Search struct {
	Group  string
	Method string
}

func (object Search) GetGroup() string {
	return object.Group
}

func (object Search) GetMethod() string {
	return object.Method
}

func (object Search) Process(params url.Values) response.Response {
	var err error
	query := "ЭД4М-0500"
	if params.Get("query") != "" {
		query = params.Get("query")
	}
	count := 5
	if params.Get("count") != "" {
		count, err = strconv.Atoi(params.Get("count"))
		if err != nil {
			return photo.Search{Status: 400}
		}
	}

	if count > 20 {
		count = 20
	}

	trains, countFound, err := parse.PhotoSearch(query, count, getParams(params))
	if err != nil {
		if err.Error() == "404" {
			return photo.Search{Status: 404}
		}
		return photo.Search{Status: 500}
	}

	return photo.Search{
		Status: 200,
		Found:  &countFound,
		Result: trains,
	}
}

func getParams(v url.Values) map[string]string {
	params := make(map[string]string)

	if v.Get("order") != "" {
		params["order"] = v.Get("order")
	}
	if v.Get("st") != "" {
		params["st"] = v.Get("st")
	}
	return params
}
