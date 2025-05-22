package models

type Route struct {
    ID        string `json:"_id,omitempty"`
    Rev       string  `json:"_rev,omitempty"` 
    Title      string  `json:"title"`
	Description string `json:"description"`
	Distance string `json:"distance"`
	Duration string `json:"duration"`
	Difficulty string `json:"difficulty"`
	ImageUrl string `json:"imageUrl"`
	Highlights string `json:"highlights"`
	MapUrl string `json:"mapUrl"`
}
