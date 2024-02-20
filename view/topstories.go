package view

import "github.com/sunesimonsen/htmx-hackernews/templates"

type TopStoriesRepo interface {
	GetTopStoryIds() ([]int, error)
}

type TopStoriesView struct {
	Templates templates.Renderer
	Repo      TopStoriesRepo
}

type TopStoriesViewData struct {
	Ids []int
}

func (v TopStoriesView) Data(params Params, headers Headers) (ViewData[TopStoriesViewData], error) {
	result := ViewData[TopStoriesViewData]{
		Template: "topstories.gohtml",
	}

	ids, err := v.Repo.GetTopStoryIds()

	result.Data.Ids = ids

	return result, err
}
