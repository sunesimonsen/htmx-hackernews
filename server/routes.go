package server

import "net/http"

func (s *server) setupRoutes() {
	s.router.GET("/", s.Index())
	s.router.GET("/story/:id", s.Story())
	s.router.ServeFiles("/assets/*filepath", http.Dir("server/assets"))
}
