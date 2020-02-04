package infrastructure

import "trainpix-api/object/train"

type Depot struct {
	Id        int            `json:"ID"`
	Name      string         `json:"name"`
	TrainList []*train.Train `json:"trains,omitempty"`
}
