package repo

import (
	"errors"
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Text       string
}

func (e HttpError) Error() string {
	text := fmt.Sprintf("%d %s", e.StatusCode, http.StatusText(e.StatusCode))

	if e.Text == "" {
		return text
	}

	return fmt.Sprintf("%s: %s", text, e.Text)
}

var NotFoundError = errors.New("Not found")
