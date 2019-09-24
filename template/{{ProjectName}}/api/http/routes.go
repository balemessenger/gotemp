package http

func (s *Server) setupRouter() {
	s.engine.GET("/health", s.handler.HealthCheck)
}
