package depot

import "trainpix-api/object"

type Search struct {
	Status int              `json:"status"`
	Result *[]*object.Depot `json:"result,omitempty"`
}

func (object Search) GetStatus() int {
	return object.Status
}
