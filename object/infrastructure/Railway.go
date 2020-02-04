package infrastructure

type Railway struct {
	Id        int      `json:"ID"`
	Name      string   `json:"name"`
	DepotList []*Depot `json:"depots,omitempty"`
}
