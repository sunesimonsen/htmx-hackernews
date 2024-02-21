package view

import "github.com/sunesimonsen/htmx-hackernews/templates"

type Index struct{}

func (v Index) Render(params Params, headers Headers) (Result, error) {
	return Result{
		Component: templates.Index(),
		HashKey:   "Index",
	}, nil
}
