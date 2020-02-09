package object

type Railway struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	DepotList *[]*Depot `json:"depots,omitempty"`
}
