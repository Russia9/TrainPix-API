package photo

import "trainpix-api/object"

type Get struct {
	Status int           `json:"status"`
	Result *object.Photo `json:"result,omitempty"`
}

func (object Get) GetStatus() int {
	return object.Status
}
