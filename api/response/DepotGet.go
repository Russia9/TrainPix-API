package response

import (
	"trainpix-api/object/infrastructure"
	"trainpix-api/object/train"
)

type DepotGet struct {
	Depot  *infrastructure.Depot `json:"depot"`
	Trains *[]*train.Train      `json:"trains,omitempty"`
}
