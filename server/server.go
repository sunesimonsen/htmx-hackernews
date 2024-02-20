package server

import (
	"net/http"

	"log/slog"

	"github.com/sunesimonsen/htmx-hackernews/middleware"
	"github.com/sunesimonsen/htmx-hackernews/repo"
	"github.com/sunesimonsen/htmx-hackernews/templates"
)

type server struct {
	templates templates.Renderer
	router    *http.ServeMux
	repo      repo.Host
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loggingMiddleware := middleware.Logging(slog.Default())

	handler := loggingMiddleware(s.router)

	handler.ServeHTTP(w, r)
}

type Config struct {
	RepoHost repo.Host
}

func NewServer(config Config) (*server, error) {
	router := http.NewServeMux()

	s := &server{router: router}

	s.templates = templates.NewRenderer()
	s.repo = config.RepoHost

	s.setupRoutes()

	return s, nil
}
