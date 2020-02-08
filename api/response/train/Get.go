package train

import "trainpix-api/object"

type Get struct {
	Status int           `json:"status"`
	Train  *object.Train `json:"train,omitempty"`
}

func (object Get) GetStatus() int {
	return object.Status
}
