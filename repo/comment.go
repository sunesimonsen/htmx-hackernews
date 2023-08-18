package repo

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (host *Host) GetComment(id string) (model.Comment, error) {
	comment := model.Comment{}
	err := host.LoadJson(fmt.Sprintf("/v0/item/%s.json", id), &comment)
	comment.Answers = len(comment.Kids)
	return comment, err
}
