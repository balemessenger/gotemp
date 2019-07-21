package server

type Server struct {
	isReady chan bool
}

func New() *Server {
	return &Server{
		isReady: make(chan bool),
	}
}

func (s *Server) Run() bool {
	go s.start()
	return <-s.isReady
}

func (s *Server) start() {
	// Write your entry point
	s.isReady <- true
}
