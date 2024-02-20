package view

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/model"
)

type StoryRepo interface {
	GetStory(id string) (model.Story, error)
}

type StoryView struct {
	Repo StoryRepo
}

func (v StoryView) Data(params Params, headers Headers) (ViewData[model.Story], error) {
	result := ViewData[model.Story]{
		Template: "story.gohtml",
	}

	story, err := v.Repo.GetStory(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Data = story

	result.HashKey = fmt.Sprintf(
		"descendants:%d,score:%d",
		story.Descendants,
		story.Score,
	)

	return result, err
}

type StoryWithCommentsView struct {
	Repo StoryRepo
}

func (v StoryWithCommentsView) Data(params Params, headers Headers) (ViewData[model.Story], error) {
	result := ViewData[model.Story]{
		Template: "story-with-comments.gohtml",
	}

	story, err := v.Repo.GetStory(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Data = story

	result.HashKey = fmt.Sprintf(
		"descendants:%d,score:%d",
		story.Descendants,
		story.Score,
	)

	return result, err
}
