package train

import "trainpix-api/object"

type Search struct {
	Status int              `json:"status"`
	Found  *int             `json:"found"`
	Parsed *int             `json:"parsed"`
	Result *[]*object.Train `json:"result,omitempty"`
}

func (object Search) GetStatus() int {
	return object.Status
}
