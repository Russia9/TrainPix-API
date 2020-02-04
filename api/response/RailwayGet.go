package response

import "trainpix-api/object/infrastructure"

type RailwayGet struct {
	Railway *infrastructure.Railway `json:"railway"`
}
