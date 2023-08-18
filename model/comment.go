package model

type Comment struct {
	Id      int    `json:"id"`
	By      string `json:"by"`
	Parent  int    `json:"parent"`
	Text    string `json:"text"`
	Time    int    `json:"time"`
	Kids    []int  `json:"kids"`
	Answers int
}
