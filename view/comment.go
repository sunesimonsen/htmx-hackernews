package view

import (
	"github.com/sunesimonsen/htmx-hackernews/model"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type CommentRepo interface {
	GetComment(id string) (model.Comment, error)
}

type CommentView struct {
	Templates templates.Renderer
	Repo      CommentRepo
}

func (v CommentView) Render(params Params, headers Headers, opt Options) ([]byte, error) {
	type Data struct {
		Comment          model.Comment
		IncludeLayout    bool
		ShowCommentsLink bool
	}

	comment, err := v.Repo.GetComment(params.Get("id"))

	if err != nil {
		return nil, err
	}

	return v.Templates.Render("comment.gohtml", opt.Layout, Data{
		Comment:          comment,
		IncludeLayout:    opt.IncludeLayout,
		ShowCommentsLink: !opt.IncludeLayout && comment.Answers > 0,
	})
}
