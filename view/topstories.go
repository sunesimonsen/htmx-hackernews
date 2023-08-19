package view

import "github.com/sunesimonsen/htmx-hackernews/templates"

type TopStoriesRepo interface {
	GetTopStoryIds() ([]int, error)
}

type TopStoriesView struct {
	Templates templates.Renderer
	Repo      TopStoriesRepo
}

func (v TopStoriesView) Render(params Params, headers Headers) ([]byte, error) {
	type Data struct {
		Ids []int
	}

	ids, err := v.Repo.GetTopStoryIds()

	if err != nil {
		return nil, err
	}

	return v.Templates.Render("topstories.gohtml", Data{
		Ids: ids,
	})
}
