package response

import "trainpix-api/object/train"

type TrainSearch struct {
	ResultCode int           `json:"result_code"`
	Trains     []train.Train `json:"trains"`
}
