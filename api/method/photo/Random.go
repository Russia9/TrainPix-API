package photo

import (
	"net/url"
	"trainpix-api/api/response"
	"trainpix-api/api/response/photo"
	"trainpix-api/parse"
)

type Random struct {
	Group  string
	Method string
}

func (object Random) GetGroup() string {
	return object.Group
}

func (object Random) GetMethod() string {
	return object.Method
}

func (object Random) Process(params url.Values) response.Response {
	result, train, err := parse.RandomPhotoGet()

	if err != nil {
		return photo.Random{Status: 500}
	}

	return photo.Random{
		Status: 200,
		Photo:  result,
		Train:  train,
	}
}
