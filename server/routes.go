package server

func (s *server) setupRoutes() {
	s.router.GET("/", s.Index())
	s.router.GET("/story/:id", s.Story())
}
