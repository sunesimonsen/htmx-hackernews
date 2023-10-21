package server

import (
	"net/http"

	"github.com/sunesimonsen/htmx-hackernews/view"
)

func (s *server) setupRoutes() {
	s.router.GET("/", view.WithView(
		s.templates,
		view.IndexView{},
	))
	s.router.GET("/topstories", view.WithView(
		s.templates,
		view.TopStoriesView{Repo: s.repo},
	))
	s.router.GET("/story/:id", view.WithView(
		s.templates,
		view.StoryView{Repo: s.repo},
	))
	s.router.GET("/comment/:id", view.WithView(
		s.templates,
		view.CommentView{Repo: s.repo},
	))
	s.router.ServeFiles("/assets/*filepath", http.Dir("server/assets"))
}
