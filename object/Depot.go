package object

type Depot struct {
	Id        int       `json:"ID"`
	Name      string    `json:"name"`
	TrainList *[]*Train `json:"trains,omitempty"`
}
