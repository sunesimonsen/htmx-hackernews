package view

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hhsnopek/etag"
	"github.com/julienschmidt/httprouter"
	"github.com/sunesimonsen/htmx-hackernews/repo"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type Headers interface {
	Get(name string) string
}

type Params interface {
	Get(name string) string
}

type paramsWrapper struct {
	httprouter.Params
}

func (ps paramsWrapper) Get(name string) string {
	return ps.ByName(name)
}

type Options struct {
	IncludeLayout bool
	Layout        string
}

type ViewData[T any] struct {
	Template string
	HashKey  string
	Data     T
}

type View[T any] interface {
	Data(params Params, headers Headers, opt Options) (ViewData[T], error)
}

func WithView[T any](renderer templates.Renderer, view View[T]) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		includeLayout := r.Header.Get("Hx-Request") == ""

		options := Options{
			IncludeLayout: includeLayout,
			Layout:        "content",
		}

		if includeLayout {
			options.Layout = "main"
		}

		data, err := view.Data(paramsWrapper{ps}, r.Header, options)

		httpError := &repo.HttpError{}
		if errors.As(err, httpError) {
			fmt.Println(err)
			http.Error(
				w,
				httpError.Error(),
				httpError.StatusCode,
			)
			return
		} else if err != nil {
			fmt.Println(err)
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Cache-Control", "max-age=0")

		if data.HashKey != "" {
			e := etag.Generate([]byte(data.HashKey), true)

			ifNoneMatch := r.Header.Get("If-None-Match")
			if ifNoneMatch == e {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		// render template
		renderer.Render(w, data.Template, options.Layout, data.Data)
	}
}
