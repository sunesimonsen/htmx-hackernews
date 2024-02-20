package view

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

type CommentRepo interface {
	GetComment(id string) (model.Comment, error)
}

type CommentView struct {
	Repo CommentRepo
}

func (v CommentView) Data(params Params, headers Headers, opt Options) (ViewData[model.Comment], error) {
	result := ViewData[model.Comment]{
		Template: "comment.gohtml",
	}

	comment, err := v.Repo.GetComment(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Data = comment

	result.HashKey = fmt.Sprintf("answers:%d", comment.Answers)

	return result, err
}

type CommentWithAnswersView struct {
	Repo CommentRepo
}

func (v CommentWithAnswersView) Data(params Params, headers Headers, opt Options) (ViewData[model.Comment], error) {
	result := ViewData[model.Comment]{
		Template: "comment-with-answers.gohtml",
	}

	comment, err := v.Repo.GetComment(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Data = comment

	result.HashKey = fmt.Sprintf("answers:%d", comment.Answers)

	return result, err
}
