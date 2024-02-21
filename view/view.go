package view

import (
	"github.com/a-h/templ"
)

type Headers interface {
	Get(name string) string
}

type Params interface {
	Get(name string) string
}

type Result struct {
	HashKey   string
	Component templ.Component
}

type View interface {
	Render(params Params, headers Headers) (Result, error)
}
