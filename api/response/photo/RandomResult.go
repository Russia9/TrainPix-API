package photo

import "trainpix-api/object"

type RandomResult struct {
	Photo *object.Photo `json:"photo,omitempty"`
	Train *object.Train `json:"train,omitempty"`
}
