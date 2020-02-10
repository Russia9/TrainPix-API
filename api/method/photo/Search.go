package photo

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/photo"
	"trainpix-api/object"
	"trainpix-api/parse"
)

type Search struct {
	Group  string
	Method string
}

func (method Search) GetGroup() string {
	return method.Group
}

func (method Search) GetMethod() string {
	return method.Method
}

func (method Search) Process(params url.Values) response.Response {
	var err error
	query := "ЭР2-1017"
	var result []*object.Photo
	var countFound int
	count := 5
	quick := false

	if params.Get("query") != "" {
		query = params.Get("query")
	}

	if params.Get("count") != "" {
		count, err = strconv.Atoi(params.Get("count"))
		if err != nil {
			return photo.Search{Status: 400}
		}
	}

	if params.Get("quick") != "" {
		quickInt, err := strconv.Atoi(params.Get("quick"))
		if err != nil {
			return photo.Search{Status: 400}
		}
		if quickInt == 1 {
			quick = true
		}
	}

	if quick {
		if count > 100 {
			count = 100
		}
	} else {
		if count > 5 {
			count = 5
		}
	}

	result, countFound, err = parse.PhotoSearch(query, count, quick, getParams(params))
	if err != nil {
		if err.Error() == "404" {
			return photo.Search{Status: 404}
		} else {
			return photo.Search{Status: 500}
		}
	}

	return photo.Search{
		Status: 200,
		Found:  &countFound,
		Result: &result,
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
