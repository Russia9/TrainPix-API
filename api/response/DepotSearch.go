package response

import "trainpix-api/object/infrastructure"

type DepotSearch struct {
	Found     int                     `json:"found"`
	DepotList []*infrastructure.Depot `json:"depots"`
}
