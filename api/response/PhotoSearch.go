package response

import "trainpix-api/object/photo"

type PhotoSearch struct {
	ResultCode int           `json:"result_code"`
	Photos     []photo.Photo `json:"photos"`
}
