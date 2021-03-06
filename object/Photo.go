package object

type Photo struct {
	ID              int     `json:"ID"`
	Image           string  `json:"image"`
	Thumbnail       string  `json:"thumbnail"`
	Page            string  `json:"page,omitempty"`
	Date            *string `json:"date,omitempty"`
	PublicationTime *string `json:"publication_time,omitempty"`
	License         *string `json:"license,omitempty"`
	Location        *string `json:"location,omitempty"`
	Author          *string `json:"author,omitempty"`
	AuthorLink      *string `json:"author_link,omitempty"`
}
