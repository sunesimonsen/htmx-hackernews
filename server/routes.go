package server

import (
	"net/http"

	"github.com/sunesimonsen/htmx-hackernews/view"
)

func (s *server) setupRoutes() {
	s.router.GET("/", view.WithView(
		view.IndexView{Templates: s.templates},
	))
	s.router.GET("/topstories", view.WithView(
		view.TopStoriesView{Templates: s.templates, Repo: s.repo},
	))
	s.router.GET("/story/:id", view.WithView(
		view.StoryView{Templates: s.templates, Repo: s.repo},
	))
	s.router.GET("/comment/:id", view.WithView(
		view.CommentView{Templates: s.templates, Repo: s.repo},
	))
	s.router.ServeFiles("/assets/*filepath", http.Dir("server/assets"))
}
