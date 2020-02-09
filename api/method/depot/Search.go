package depot

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/depot"
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
	railway := -1
	quick := false
	count := 5

	query := "ЭР2"
	if params.Get("query") != "" {
		query = params.Get("query")
	}

	if params.Get("railway") != "" {
		railway, err = strconv.Atoi(params.Get("railway"))

		if err != nil {
			return depot.Search{Status: 400}
		}
	} else {
		return depot.Search{Status: 400}
	}

	if params.Get("count") != "" {
		count, err = strconv.Atoi(params.Get("count"))
		if err != nil {
			return depot.Search{Status: 400}
		}
	}

	if count > 10 {
		count = 10
	}

	if params.Get("quick") == "1" {
		quick = true
	}

	trains := 5
	if params.Get("trains") != "" {
		trains, err = strconv.Atoi(params.Get("trains"))
		if err != nil {
			return depot.Search{Status: 400}
		}
	}

	if quick && trains > 100 {
		trains = 100
	} else if trains > 10 {
		trains = 10
	}

	depots, err := parse.DepotSearch(query, railway, count, trains, quick)

	if err != nil {
		if err.Error() == "404" {
			return depot.Search{Status: 404}
		}
		return depot.Search{Status: 500}
	}

	return depot.Search{
		Status: 200,
		Result: &depots,
	}
}
