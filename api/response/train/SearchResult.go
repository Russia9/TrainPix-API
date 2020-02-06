package train

import "trainpix-api/object"

type SearchResult struct {
	CountFound  int             `json:"found,omitempty"`
	CountParsed int             `json:"parsed,omitempty"`
	Result      []*object.Train `json:"result,omitempty"`
}
