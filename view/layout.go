package view

import (
	"github.com/a-h/templ"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type Layout func(part templ.Component) templ.Component

type layoutedView struct {
	view   View
	layout Layout
}

func (lv layoutedView) Render(params Params, headers Headers) (Result, error) {
	result, err := lv.view.Render(params, headers)

	if err != nil {
		return result, err
	}

	return Result{
		Component: lv.layout(result.Component),
		HashKey:   result.HashKey,
	}, nil
}

func WithLayout(layout Layout, view View) View {
	return layoutedView{layout: layout, view: view}
}

func WithMainLayout(view View) View {
	return WithLayout(templates.MainLayout, view)
}
