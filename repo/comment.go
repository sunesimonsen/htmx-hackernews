package repo

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (host *Host) GetComment(id string) (model.Comment, error) {
	comment := model.Comment{}
	err := host.LoadJson(fmt.Sprintf("/v0/item/%s.json", id), &comment)
	comment.Answers = len(comment.Kids)

	// zero is not a valid id
	if comment.Id == 0 {
		return comment, fmt.Errorf("%w: comment %s", NotFoundError, id)
	}

	return comment, err
}
