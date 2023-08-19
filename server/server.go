package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sunesimonsen/htmx-hackernews/repo"
	"github.com/sunesimonsen/htmx-hackernews/view"
)

type server struct {
	templates view.TemplateRenderer
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

	s.templates = view.NewTemplateRenderer()
	s.repo = config.RepoHost

	s.setupRoutes()

	return s, nil
}
