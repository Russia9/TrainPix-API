package train

import (
	"errors"
	"net/url"
	"strconv"
	"trainpix-api/object"
	"trainpix-api/parser"
)

func Get(query url.Values) (*object.Train, error) {
	if query.Get("id") == "" {
		return nil, errors.New("400")
	}
	id, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		return nil, errors.New("400")
	}

	quick := false
	if query.Get("quick") == "1" {
		quick = true
	}

	result, err := parser.TrainGet(id, quick)

	return result, err
}
