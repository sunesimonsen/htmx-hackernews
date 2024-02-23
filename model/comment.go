package model

import "github.com/kennygrant/sanitize"

type Comment struct {
	Id     int    `json:"id"`
	By     string `json:"by"`
	Parent int    `json:"parent"`
	Text   string `json:"text"`
	Time   int    `json:"time"`
	Kids   []int  `json:"kids"`
}

func (c Comment) Html() string {
	html, err := sanitize.HTMLAllowing(
		c.Text,
		[]string{"p", "a", "strong", "em"},
	)

	if err != nil {
		sanitize.HTML(c.Text)
	}

	return html
}

func (c Comment) Answers() int {
	return len(c.Kids)
}
