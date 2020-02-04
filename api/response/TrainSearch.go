package response

import "trainpix-api/object/train"

type TrainSearch struct {
	Found     int            `json:"found"`
	Parsed    int            `json:"parsed"`
	TrainList []*train.Train `json:"trains"`
}
