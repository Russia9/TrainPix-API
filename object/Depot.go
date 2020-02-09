package object

type Depot struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	TrainList *[]*Train `json:"trains,omitempty"`
}
