package train

import (
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/train"
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
	query := "ЭР2"
	var result []*object.Train
	var countFound int
	var countParsed int
	count := 5
	photoCount := 1
	quick := false
	photoQuick := true

	if params.Get("query") != "" {
		query = params.Get("query")
	}

	if params.Get("count") != "" {
		count, err = strconv.Atoi(params.Get("count"))
		if err != nil {
			return train.Search{Status: 400}
		}
	}

	if params.Get("photoCount") != "" {
		photoCount, err = strconv.Atoi(params.Get("photoCount"))
		if err != nil {
			return train.Search{Status: 400}
		}
	}

	if params.Get("quick") != "" {
		quickInt, err := strconv.Atoi(params.Get("quick"))
		if err != nil {
			return train.Search{Status: 400}
		}
		if quickInt == 1 {
			quick = true
		}
	}

	if params.Get("photoQuick") != "" {
		quickInt, err := strconv.Atoi(params.Get("photoQuick"))
		if err != nil {
			return train.Get{Status: 400}
		}
		if quickInt == 0 {
			photoQuick = false
		}
	}

	if photoCount > 10 {
		photoCount = 10
	}

	if quick {
		if count > 100 {
			count = 100
		}
	} else {
		if count > 20 {
			count = 20
		}
	}

	result, countFound, countParsed, err = parse.TrainSearch(query, count, photoCount, quick, photoQuick, getParams(params))
	if err != nil {
		if err.Error() == "404" {
			return train.Get{Status: 404}
		} else {
			return train.Get{Status: 500}
		}
	}

	return train.Search{
		Status: 200,
		Found:  &countFound,
		Parsed: &countParsed,
		Result: &result,
	}
}

func getParams(v url.Values) map[string]string {
	params := make(map[string]string)

	if v.Get("state") != "" {
		params["state"] = v.Get("state")
	}
	if v.Get("order") != "" {
		params["order"] = v.Get("order")
	}
	if v.Get("st") != "" {
		params["st"] = v.Get("st")
	}
	return params
}
