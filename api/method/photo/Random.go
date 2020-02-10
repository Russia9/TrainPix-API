package photo

import (
	"net/url"
	"trainpix-api/api/response"
	"trainpix-api/api/response/photo"
	"trainpix-api/object"
	"trainpix-api/parse"
)

type Random struct {
	Group  string
	Method string
}

func (method Random) GetGroup() string {
	return method.Group
}

func (method Random) GetMethod() string {
	return method.Method
}

func (method Random) Process(_ url.Values) response.Response {
	var err error
	var photoObject *object.Photo
	var trainObject *object.Train

	photoObject, trainObject, err = parse.RandomPhotoGet()
	if err != nil {
		if err.Error() == "404" {
			return photo.Random{Status: 404}
		} else {
			return photo.Random{Status: 500}
		}
	}

	return photo.Random{
		Status: 200,
		Photo:  photoObject,
		Train:  trainObject,
	}
}
