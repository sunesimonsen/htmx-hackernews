package server

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sunesimonsen/htmx-hackernews/repo"
)

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params) error

type BufferedResponseWriter struct {
	wrapped    http.ResponseWriter
	header     http.Header
	body       bytes.Buffer
	statusCode int
}

func WithErrorHandling(handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := handle(w, r, ps)
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
	}
}
