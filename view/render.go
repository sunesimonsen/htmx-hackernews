package view

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hhsnopek/etag"
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
	req *http.Request
}

func (ps paramsWrapper) Get(name string) string {
	return ps.req.PathValue(name)
}

type ViewData[T any] struct {
	Template string
	HashKey  string
	Data     T
}

type View[T any] interface {
	Data(params Params, headers Headers) (ViewData[T], error)
}

func WithView[T any](renderer templates.Renderer, layout string, view View[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := view.Data(paramsWrapper{req: r}, r.Header)

		httpError := &repo.HttpError{}
		if errors.As(err, httpError) {
			http.Error(
				w,
				httpError.Error(),
				httpError.StatusCode,
			)
			return
		}

		if errors.Is(err, repo.NotFoundError) {
			http.Error(
				w,
				err.Error(),
				http.StatusNotFound,
			)
			return
		}

		if err != nil {
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
			hashKey := fmt.Sprintf("layout:%s,%s", layout, data.HashKey)
			e := etag.Generate([]byte(hashKey), true)

			ifNoneMatch := r.Header.Get("If-None-Match")
			if ifNoneMatch == e {
				w.WriteHeader(http.StatusNotModified)
				return
			}

			w.Header().Set("ETag", e)
		}

		// render template
		renderer.Render(w, data.Template, layout, data.Data)
	}
}
