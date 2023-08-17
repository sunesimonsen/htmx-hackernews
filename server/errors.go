package server

import (
	"bytes"
	"errors"
	"fmt"
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

type BufferedResponseWriter struct {
	wrapped    http.ResponseWriter
	header     http.Header
	body       bytes.Buffer
	statusCode int
}

func WithErrorHandling(handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := handle(w, r, ps)
		httpError := &HttpError{}
		if errors.As(err, httpError) {
			fmt.Println(err)
			http.Error(
				w,
				httpError.Error(),
				httpError.Code,
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
