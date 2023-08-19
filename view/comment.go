package view

import "github.com/sunesimonsen/htmx-hackernews/model"

type CommentRepo interface {
	GetComment(id string) (model.Comment, error)
}

type CommentView struct {
	Templates TemplateRenderer
	Repo      CommentRepo
}

func (v CommentView) Render(params Params, headers Headers) ([]byte, error) {
	type Data struct {
		Comment          model.Comment
		ShowCommentsLink bool
		IncludeLayout    bool
	}

	comment, err := v.Repo.GetComment(params.Get("id"))

	if err != nil {
		return nil, err
	}

	includeLayout := headers.Get("Hx-Request") == ""

	return v.Templates.Render("comment.gohtml", Data{
		Comment:          comment,
		IncludeLayout:    includeLayout,
		ShowCommentsLink: !includeLayout && comment.Answers > 0,
	})
}
