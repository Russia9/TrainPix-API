package photo

import "trainpix-api/object"

type Random struct {
	Status int           `json:"status"`
	Photo  *object.Photo `json:"photo,omitempty"`
	Train  *object.Train `json:"train,omitempty"`
}

func (object Random) GetStatus() int {
	return object.Status
}
