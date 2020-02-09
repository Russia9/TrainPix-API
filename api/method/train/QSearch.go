package train

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/train"
	"trainpix-api/parse"
)

type QSearch struct {
	Group  string
	Method string
}

func (object QSearch) GetGroup() string {
	return object.Group
}

func (object QSearch) GetMethod() string {
	return object.Method
}

func (object QSearch) Process(params url.Values) response.Response {
	var err error
	query := "ЭР2"
	if params.Get("query") != "" {
		query = params.Get("query")
	}
	count := 5
	if params.Get("count") != "" {
		count, err = strconv.Atoi(params.Get("count"))
		if err != nil {
			return train.Search{Status: 400}
		}
	}

	if count > 100 {
		count = 100
	}

	trains, countFound, countParsed, err := parse.TrainSearch(query, count, true, getParams(params))
	if err != nil {
		if err.Error() == "404" {
			return train.Search{Status: 404}
		}
		return train.Search{Status: 500}
	}

	return train.Search{
		Status: 200,
		Found:  &countFound,
		Parsed: &countParsed,
		Result: &trains,
	}
}
