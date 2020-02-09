package object

type Train struct {
	ID                   int       `json:"ID"`
	Name                 string    `json:"name"`
	Railway              *Railway  `json:"railway,omitempty"`
	Depot                *Depot    `json:"depot,omitempty"`
	Model                *Model    `json:"model,omitempty"`
	Builder              *string   `json:"builder,omitempty"`
	IdentificationNumber *string   `json:"identification_number,omitempty"`
	SerialType           *string   `json:"serial_type,omitempty"`
	Built                *string   `json:"built,omitempty"`
	Category             *string   `json:"category,omitempty"`
	Condition            int       `json:"condition,omitempty"`
	Note                 *string   `json:"note,omitempty"`
	Info                 *string   `json:"info,omitempty"`
	PhotoList            *[]*Photo `json:"photos,omitempty"`
}
