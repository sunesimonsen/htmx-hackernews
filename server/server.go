package server

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/sunesimonsen/htmx-hackernews/repo"
)

type server struct {
	templates *template.Template
	router    *httprouter.Router
	repo      repo.Host
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

type Config struct {
	RepoHost repo.Host
}

func NewServer(config Config) (*server, error) {
	router := httprouter.New()

	s := &server{router: router}
	err := s.setupTemplates()
	if err != nil {
		return nil, err
	}

	s.repo = config.RepoHost
	s.setupRoutes()

	return s, nil
}
