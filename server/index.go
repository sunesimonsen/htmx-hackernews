package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *server) Index() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := s.templates.ExecuteTemplate(w, "index.gohtml", nil)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
