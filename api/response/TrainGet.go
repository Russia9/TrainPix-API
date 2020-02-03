package response

import "trainpix-api/object/train"

type TrainGet struct {
	Train *train.Train `json:"train"`
}
