package response

import (
	"trainpix-api/object/photo"
	"trainpix-api/object/train"
)

type PhotoRandomGet struct {
	ResultCode int          `json:"result_code"`
	Photo      *photo.Photo `json:"photo"`
	Train      *train.Train `json:"train"`
}
