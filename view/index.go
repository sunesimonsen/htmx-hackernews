package view

import (
	"fmt"

	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type TopStoriesRepo interface {
	GetTopStoryIds() ([]int, error)
}

type Index struct {
	Repo TopStoriesRepo
}

func (v Index) Render(params Params, headers Headers) (Result, error) {
	var result Result

	ids, err := v.Repo.GetTopStoryIds()

	result.Component = templates.Index(ids)
	result.HashKey = fmt.Sprintln(ids)

	return result, err
}
