package server

import (
	"net/http"

	"log/slog"

	"github.com/sunesimonsen/htmx-hackernews/middleware"
	"github.com/sunesimonsen/htmx-hackernews/repo"
)

type server struct {
	router Router
	repo   repo.Host
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stack := middleware.CreateStack(
		middleware.Logging(slog.Default()),
	)

	handler := stack(s.router)

	handler.ServeHTTP(w, r)
}

type Config struct {
	RepoHost repo.Host
}

func NewServer(config Config) (*server, error) {
	router := NewRouter()

	s := &server{router: router}

	s.repo = config.RepoHost

	s.setupRoutes()

	return s, nil
}
