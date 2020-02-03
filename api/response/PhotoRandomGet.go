package response

import (
	"trainpix-api/object/photo"
	"trainpix-api/object/train"
)

type PhotoRandomGet struct {
	Photo *photo.Photo `json:"photo"`
	Train *train.Train `json:"train"`
}
