package photo

type Photo struct {
	Id         int         `json:"ID"`
	Image      string      `json:"image"`
	Thumbnail  string      `json:"thumbnail"`
	Page       string      `json:"page,omitempty"`
	Date       *string     `json:"date,omitempty"`
	Location   *string     `json:"location,omitempty"`
	Author     *string     `json:"author,omitempty"`
	AuthorLink *string     `json:"author_link,omitempty"`
}
