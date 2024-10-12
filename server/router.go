package server

import (
	"errors"
	"net/http"

	"github.com/hhsnopek/etag"
	"github.com/sunesimonsen/htmx-hackernews/repo"
	"github.com/sunesimonsen/htmx-hackernews/view"
)

type Router struct {
	mux *http.ServeMux
}

func (router Router) Handle(pattern string, handler http.Handler) {
	router.mux.Handle(pattern, handler)
}

func (router Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	router.mux.HandleFunc(pattern, handler)
}

func NewRouter() Router {
	return Router{
		mux: http.NewServeMux(),
	}
}

type paramsWrapper struct {
	req *http.Request
}

func (ps paramsWrapper) Get(name string) string {
	return ps.req.PathValue(name)
}

func (router Router) RegisterView(pattern string, view view.View) {
	router.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		viewResult, err := view.Render(paramsWrapper{r}, r.Header)

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

		if viewResult.HashKey != "" {
			e := etag.Generate([]byte(viewResult.HashKey), true)

			ifNoneMatch := r.Header.Get("If-None-Match")
			if ifNoneMatch == e {
				w.WriteHeader(http.StatusNotModified)
				return
			}

			w.Header().Set("ETag", e)
		}

		viewResult.Component.Render(r.Context(), w)
	})
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}
