package model

type Story struct {
	Id          int    `json:"id"`
	By          string `json:"by"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Descendants int    `json:"descendants"`
	Time        int    `json:"time"`
	Score       int    `json:"score"`
	Kids        []int  `json:"kids"`
}
