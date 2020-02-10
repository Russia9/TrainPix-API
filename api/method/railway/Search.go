package railway

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/railway"
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
	query := "Московская"
	var result []*object.Railway
	var countFound int
	depotCount := 5
	trainCount := 5
	photoCount := 5
	quick := false

	if params.Get("query") != "" {
		query = params.Get("query")
	}

	if params.Get("depotCount") != "" {
		depotCount, err = strconv.Atoi(params.Get("depotCount"))
		if err != nil {
			return railway.Get{Status: 400}
		}
	}

	if params.Get("trainCount") != "" {
		trainCount, err = strconv.Atoi(params.Get("trainCount"))
		if err != nil {
			return railway.Get{Status: 400}
		}
	}

	if params.Get("photoCount") != "" {
		photoCount, err = strconv.Atoi(params.Get("photoCount"))
		if err != nil {
			return railway.Get{Status: 400}
		}
	}

	if params.Get("quick") != "" {
		quickInt, err := strconv.Atoi(params.Get("quick"))
		if err != nil {
			return railway.Search{Status: 400}
		}
		if quickInt == 1 {
			quick = true
		}
	}

	result, countFound, err = parse.RailwaySearch(query, depotCount, trainCount, photoCount, quick)
	if err != nil {
		if err.Error() == "404" {
			return railway.Get{Status: 404}
		} else {
			return railway.Get{Status: 500}
		}
	}

	return railway.Search{
		Status: 200,
		Found:  &countFound,
		Result: &result,
	}
}
