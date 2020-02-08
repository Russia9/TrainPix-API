package photo

import "trainpix-api/object"

type Search struct {
	Status int              `json:"status"`
	Found  *int             `json:"found,omitempty"`
	Result *[]*object.Photo `json:"result,omitempty"`
}

func (object Search) GetStatus() int {
	return object.Status
}
