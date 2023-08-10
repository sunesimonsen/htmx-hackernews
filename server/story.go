package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Story struct {
	Id string
}

func (s *server) Story() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		story := Story{Id: ps.ByName("id")}

		err := s.templates.ExecuteTemplate(
			w,
			"story.gohtml",
			story,
		)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
