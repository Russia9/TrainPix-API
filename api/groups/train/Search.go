package train

import (
	"errors"
	"net/url"
	"strconv"
	"trainpix-api/api/response/train"
	"trainpix-api/parser"
)

func Search(query url.Values, quick bool) (*train.SearchResult, error) {
	var err error

	searchQuery := query.Get("query")

	count := 5
	if query.Get("count") != "" {
		count, err = strconv.Atoi(query.Get("count"))
		if err != nil {
			return nil, errors.New("400")
		}
	}
	if quick && count > 100 {
		count = 100
	} else if count > 20 {
		count = 20
	}

	result, found, parsed, err := parser.TrainSearch(searchQuery, count, quick, getParams(query))

	return &train.SearchResult{
		CountFound:  found,
		CountParsed: parsed,
		Result:      result,
	}, err
}

func getParams(query url.Values) map[string]string {
	params := make(map[string]string)

	if query.Get("state") != "" {
		params["state"] = query.Get("state")
	}
	if query.Get("order") != "" {
		params["order"] = query.Get("order")
	}
	if query.Get("st") != "" {
		params["st"] = query.Get("st")
	}
	return params
}
