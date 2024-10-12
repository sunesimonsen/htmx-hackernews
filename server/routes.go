package server

import (
	"net/http"

	"github.com/sunesimonsen/htmx-hackernews/view"
)

func (s *server) setupRoutes() {
	s.router.RegisterView("GET /{$}", view.WithMainLayout(view.Index{Repo: s.repo}))
	s.router.RegisterView("GET /story/{id}", view.WithMainLayout(view.Story{Repo: s.repo, IncludeComments: true}))
	s.router.RegisterView("GET /comment/{id}", view.WithMainLayout(view.Comment{Repo: s.repo, IncludeAnswers: true}))

	s.router.RegisterView("GET /parts/story/{id}", view.Story{Repo: s.repo})
	s.router.RegisterView("GET /parts/comment/{id}", view.Comment{Repo: s.repo})

	s.router.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./server/assets"))))

	s.router.HandleFunc("GET /up", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}
