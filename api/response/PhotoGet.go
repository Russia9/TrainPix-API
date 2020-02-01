package response

import "trainpix-api/object/photo"

type PhotoGet struct {
	ResultCode int         `json:"result_code"`
	Photo      *photo.Photo `json:"photo"`
}
