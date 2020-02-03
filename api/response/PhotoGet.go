package response

import "trainpix-api/object/photo"

type PhotoGet struct {
	Photo *photo.Photo `json:"photo"`
}
