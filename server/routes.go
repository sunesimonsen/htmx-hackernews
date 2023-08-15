package server

import "net/http"

func (s *server) setupRoutes() {
	s.router.GET("/", WithErrorHandling(s.Index()))
	s.router.GET("/story/:id", WithErrorHandling(s.Story()))
	s.router.ServeFiles("/assets/*filepath", http.Dir("server/assets"))
}
