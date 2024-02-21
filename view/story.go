package view

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type StoryRepo interface {
	GetStory(id string) (model.Story, error)
}

type Story struct {
	Repo            StoryRepo
	IncludeComments bool
}

func (v Story) Render(params Params, headers Headers) (Result, error) {
	var result Result

	story, err := v.Repo.GetStory(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Component = templates.Story(story, v.IncludeComments)

	result.HashKey = fmt.Sprintf(
		"descendants:%d,score:%d",
		story.Descendants,
		story.Score,
	)

	return result, err
}
