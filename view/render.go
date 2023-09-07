package view

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hhsnopek/etag"
	"github.com/julienschmidt/httprouter"
	"github.com/sunesimonsen/htmx-hackernews/repo"
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

type View interface {
	Render(params Params, headers Headers, opt Options) ([]byte, error)
}

func WithView(view View) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		includeLayout := r.Header.Get("Hx-Request") == ""

		options := Options{
			IncludeLayout: includeLayout,
			Layout:        "content",
		}

		if includeLayout {
			options.Layout = "main"
		}

		data, err := view.Render(paramsWrapper{ps}, r.Header, options)

		httpError := &repo.HttpError{}
		if errors.As(err, httpError) {
			fmt.Println(err)
			http.Error(
				w,
				httpError.Error(),
				httpError.StatusCode,
			)
		} else if err != nil {
			fmt.Println(err)
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}

		e := etag.Generate(data, true)

		ifNoneMatch := r.Header.Get("If-None-Match")
		if ifNoneMatch == e {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Etag", e)
		w.Header().Set("Cache-Control", "max-age=0")

		w.Write(data)
	}
}
