package railway

import "trainpix-api/object"

type Search struct {
	Status int                `json:"status"`
	Found  *int               `json:"found,omitempty"`
	Result *[]*object.Railway `json:"result,omitempty"`
}

func (object Search) GetStatus() int {
	return object.Status
}
