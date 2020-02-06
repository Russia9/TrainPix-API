package photo

import (
	"trainpix-api/api/response/photo"
	"trainpix-api/parser"
)

func RandomGet() (photo.RandomResult, error) {
	result, train, err := parser.RandomPhotoGet()

	return photo.RandomResult{
		Photo: result,
		Train: train,
	}, err
}
