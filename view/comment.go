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

type CommentViewData struct {
	Comment          model.Comment
	IncludeLayout    bool
	ShowCommentsLink bool
}

func (v CommentView) Data(params Params, headers Headers, opt Options) (ViewData[CommentViewData], error) {
	result := ViewData[CommentViewData]{
		Template: "comment.gohtml",
	}

	comment, err := v.Repo.GetComment(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Data = CommentViewData{
		Comment:          comment,
		IncludeLayout:    opt.IncludeLayout,
		ShowCommentsLink: !opt.IncludeLayout && comment.Answers > 0,
	}

	result.HashKey = fmt.Sprintf("answers:%d", comment.Answers)

	return result, err
}
