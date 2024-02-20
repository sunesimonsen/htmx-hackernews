package server

import (
	"net/http"

	"github.com/sunesimonsen/htmx-hackernews/view"
)

func (s *server) setupRoutes() {
	s.router.Handle("GET /{$}",
		view.WithView(s.templates, "main", view.IndexView{}),
	)
	s.router.Handle("GET /story/{id}",
		view.WithView(s.templates, "main", view.StoryWithCommentsView{Repo: s.repo}),
	)
	s.router.Handle("GET /comment/{id}",
		view.WithView(s.templates, "main", view.CommentWithAnswersView{Repo: s.repo}),
	)

	s.router.Handle("GET /parts/topstories",
		view.WithView(s.templates, "part", view.TopStoriesView{Repo: s.repo}),
	)
	s.router.Handle("GET /parts/story/{id}",
		view.WithView(s.templates, "part", view.StoryView{Repo: s.repo}),
	)
	s.router.Handle("GET /parts/comment/{id}",
		view.WithView(s.templates, "part", view.CommentView{Repo: s.repo}),
	)

	s.router.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./server/assets"))))
}
