package train

import (
	"errors"
	"net/url"
	"strconv"
	"trainpix-api/api/response"
	"trainpix-api/api/response/train"
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

func (object Get) Process(params url.Values) (response.Response, error) {
	id := -1
	var err error

	if params.Get("id") != "" {
		id, err = strconv.Atoi(params.Get("id"))

		if err != nil {
			return train.Get{Status: 400}, errors.New("400")
		}
	} else {
		return train.Get{Status: 400}, errors.New("400")
	}

	result, err := parse.TrainGet(id, false)

	if err != nil {
		if err.Error() == "404" {
			return train.Get{Status: 404}, errors.New("404")
		} else {
			return train.Get{Status: 500}, errors.New("500")
		}
	}

	return train.Get{Status: 200, Train: result}, nil
}
