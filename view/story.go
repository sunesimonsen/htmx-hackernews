package view

import "github.com/sunesimonsen/htmx-hackernews/model"

type StoryRepo interface {
	GetStory(id string) (model.Story, error)
}

type StoryView struct {
	Templates TemplateRenderer
	Repo      StoryRepo
}

func (v StoryView) Render(params Params, headers Headers) ([]byte, error) {
	type Data struct {
		Story            model.Story
		IncludeLayout    bool
		ShowCommentsLink bool
	}

	story, err := v.Repo.GetStory(params.Get("id"))

	if err != nil {
		return nil, err
	}

	includeLayout := headers.Get("Hx-Request") == ""

	return v.Templates.Render("story.gohtml", Data{
		Story:            story,
		IncludeLayout:    includeLayout,
		ShowCommentsLink: !includeLayout && len(story.Kids) > 0,
	})
}
