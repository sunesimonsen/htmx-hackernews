package repo

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

func (host Host) GetComment(id string) (model.Comment, error) {
	comment := model.Comment{}
	err := host.LoadJson(fmt.Sprintf("/v0/item/%s.json", id), &comment)
	if err != nil {
		return comment, err
	}

	// zero is not a valid id
	if comment.Id == 0 {
		return comment, fmt.Errorf("%w: comment %s", NotFoundError, id)
	}

	comment.Answers = len(comment.Kids)

	return comment, nil
}
