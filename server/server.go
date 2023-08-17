package server

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	templates *template.Template
	router    *httprouter.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer() (*server, error) {
	router := httprouter.New()

	s := &server{router: router}
	err := s.setupTemplates()
	if err != nil {
		return nil, err
	}

	s.setupRoutes()

	return s, nil
}
