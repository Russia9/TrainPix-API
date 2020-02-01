package train

import (
	"trainpix-api/object/infrastructure"
	"trainpix-api/object/photo"
)

type Train struct {
	Id                   int                     `json:"ID"`
	Name                 string                  `json:"name"`
	Railway              infrastructure.Railway `json:"railway"`
	Depot                infrastructure.Depot   `json:"depot"`
	Model                Model                  `json:"model"`
	Builder              *string                 `json:"builder"`
	IdentificationNumber *string                 `json:"identification_number"`
	SerialType           *string                 `json:"serial_type"`
	Built                *string                 `json:"built"`
	Category             *string                 `json:"category"`
	Condition            int                     `json:"condition"`
	Note                 *string                 `json:"note"`
	PhotoList            []*photo.Photo          `json:"photos"`
}
