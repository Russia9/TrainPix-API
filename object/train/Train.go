package train

import (
	"trainpix-api/object/infrastructure"
	"trainpix-api/object/photo"
)

type Train struct {
	Id                   int                     `json:"ID"`
	Name                 string                  `json:"name"`
	Railway              *infrastructure.Railway `json:"railway,omitempty"`
	Depot                *infrastructure.Depot   `json:"depot,omitempty"`
	Model                *Model                  `json:"model,omitempty"`
	Builder              *string                 `json:"builder,omitempty"`
	IdentificationNumber *string                 `json:"identification_number,omitempty"`
	SerialType           *string                 `json:"serial_type,omitempty"`
	Built                *string                 `json:"built,omitempty"`
	Category             *string                 `json:"category,omitempty"`
	Condition            int                     `json:"condition,omitempty"`
	Note                 *string                 `json:"note,omitempty"`
	Info                 *string                 `json:"info,omitempty"`
	PhotoList            []*photo.Photo          `json:"photos,omitempty"`
}
