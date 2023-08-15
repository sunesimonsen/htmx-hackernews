package server

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params) error

type HttpError struct {
	Code int
	Text string
}

func (e HttpError) Error() string {
	return http.StatusText(e.Code) + ": " + e.Text
}

func WithErrorHandling(handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := handle(w, r, ps)
		httpError := &HttpError{}
		if errors.As(err, httpError) {
			http.Error(
				w,
				httpError.Error(),
				httpError.Code,
			)
		} else if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}
	}
}
