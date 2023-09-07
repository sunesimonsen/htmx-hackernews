package view

import (
	"github.com/sunesimonsen/htmx-hackernews/model"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type StoryRepo interface {
	GetStory(id string) (model.Story, error)
}

type StoryView struct {
	Templates templates.Renderer
	Repo      StoryRepo
}

func (v StoryView) Render(params Params, headers Headers, opt Options) ([]byte, error) {
	type Data struct {
		Story            model.Story
		IncludeLayout    bool
		ShowCommentsLink bool
	}

	story, err := v.Repo.GetStory(params.Get("id"))

	if err != nil {
		return nil, err
	}

	return v.Templates.Render("story.gohtml", opt.Layout, Data{
		Story:            story,
		IncludeLayout:    opt.IncludeLayout,
		ShowCommentsLink: !opt.IncludeLayout && len(story.Kids) > 0,
	})
}
