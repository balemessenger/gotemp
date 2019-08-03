package http

func (s *Server) setupRouter() {
	s.engine.GET("/health", s.handler.HealthCheck)

	s.engine.POST("/example", s.handler.Example)
	s.authorized.POST("/example", s.handler.AdminExample)

}
