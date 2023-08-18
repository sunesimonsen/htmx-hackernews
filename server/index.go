package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) Index() Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
		return s.renderTemplate(w, r, "index.gohtml", nil)
	}
}
