package response

import "trainpix-api/object/train"

type TrainGet struct {
	ResultCode int         `json:"result_code"`
	Train      train.Train `json:"train"`
}
