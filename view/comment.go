package view

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type CommentRepo interface {
	GetComment(id string) (model.Comment, error)
}

type Comment struct {
	Repo           CommentRepo
	IncludeAnswers bool
}

func (v Comment) Render(params Params, headers Headers) (Result, error) {
	var result Result

	comment, err := v.Repo.GetComment(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Component = templates.Comment(comment, v.IncludeAnswers)

	result.HashKey = fmt.Sprintf("answers:%d", comment.Answers)

	return result, err
}
