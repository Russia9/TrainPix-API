package depot

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/depot"
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
	var query string
	var result []*object.Depot
	var railwayID int
	count := 5
	trainCount := 5
	photoCount := 5
	quick := true

	if params.Get("query") != "" {
		query = params.Get("query")
	}

	if params.Get("railway") == "" {
		return depot.Search{Status: 400}
	} else {
		railwayID, err = strconv.Atoi(params.Get("railway"))
		if err != nil {
			return depot.Search{Status: 400}
		}
	}

	if params.Get("count") != "" {
		count, err = strconv.Atoi(params.Get("count"))
		if err != nil {
			return depot.Search{Status: 400}
		}
	}

	if params.Get("trainCount") != "" {
		trainCount, err = strconv.Atoi(params.Get("trainCount"))
		if err != nil {
			return depot.Search{Status: 400}
		}
	}

	if params.Get("photoCount") != "" {
		photoCount, err = strconv.Atoi(params.Get("photoCount"))
		if err != nil {
			return depot.Search{Status: 400}
		}
	}

	if params.Get("quick") != "" {
		quickInt, err := strconv.Atoi(params.Get("quick"))
		if err != nil {
			return depot.Search{Status: 400}
		}
		if quickInt == 0 {
			quick = false
		}
	}

	result, err = parse.DepotSearch(query, railwayID, count, trainCount, photoCount, quick)

	if err != nil {
		if err.Error() == "404" {
			return depot.Search{Status: 404}
		} else {
			return depot.Search{Status: 500}
		}
	}

	return depot.Search{
		Status: 200,
		Result: &result,
	}
}
