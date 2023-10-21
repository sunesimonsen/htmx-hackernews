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

type StoryViewData struct {
	Story            model.Story
	IncludeLayout    bool
	ShowCommentsLink bool
}

func (v StoryView) Data(params Params, headers Headers, opt Options) (ViewData[StoryViewData], error) {
	result := ViewData[StoryViewData]{
		Template: "story.gohtml",
	}

	story, err := v.Repo.GetStory(params.Get("id"))

	if err != nil {
		return result, err
	}

	result.Data = StoryViewData{
		Story:            story,
		IncludeLayout:    opt.IncludeLayout,
		ShowCommentsLink: !opt.IncludeLayout && len(story.Kids) > 0,
	}

	result.HashKey = fmt.Sprintf(
		"descendants:%d,score:%d",
		story.Descendants,
		story.Score,
	)

	return result, err
}
