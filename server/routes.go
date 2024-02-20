package server

import (
	"net/http"

	"github.com/sunesimonsen/htmx-hackernews/view"
)

func (s *server) setupRoutes() {
	s.router.GET("/",
		view.WithView(s.templates, "main", view.IndexView{}),
	)
	s.router.GET("/story/:id",
		view.WithView(s.templates, "main", view.StoryWithCommentsView{Repo: s.repo}),
	)
	s.router.GET("/comment/:id",
		view.WithView(s.templates, "main", view.CommentView{Repo: s.repo}),
	)

	s.router.GET("/parts/topstories",
		view.WithView(s.templates, "content", view.TopStoriesView{Repo: s.repo}),
	)
	s.router.GET("/parts/story/:id",
		view.WithView(s.templates, "content", view.StoryView{Repo: s.repo}),
	)
	s.router.GET("/parts/comment/:id",
		view.WithView(s.templates, "content", view.CommentView{Repo: s.repo}),
	)

	s.router.ServeFiles("/assets/*filepath", http.Dir("server/assets"))
}
