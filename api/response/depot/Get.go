package depot

import "trainpix-api/object"

type Get struct {
	Status int           `json:"status"`
	Result *object.Depot `json:"result,omitempty"`
}

func (object Get) GetStatus() int {
	return object.Status
}
