package response

import "trainpix-api/object/infrastructure"

type DepotGet struct {
	Depot *infrastructure.Depot `json:"depot"`
}
