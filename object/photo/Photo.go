package photo

type Photo struct {
	Id         int    `json:"ID"`
	Image      string `json:"image"`
	Thumbnail  string `json:"thumbnail"`
	Page       string `json:"page"`
	Date       string `json:"date"`
	Location   string `json:"location"`
	Author     string `json:"author"`
	AuthorLink string `json:"author_link"`
}
