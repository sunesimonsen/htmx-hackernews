package view

import "github.com/sunesimonsen/htmx-hackernews/templates"

type IndexView struct {
	Templates templates.Renderer
}

func (v IndexView) Render(params Params, headers Headers) ([]byte, error) {
	return v.Templates.Render("index.gohtml", nil)
}
