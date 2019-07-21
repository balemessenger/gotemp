package http

import (
	"github.com/gin-gonic/gin"
	"{{ProjectName}}/pkg"
	"sync"
)

type Server struct {
	engine     *gin.Engine
	authorized *gin.RouterGroup
	handler    *Handler
}

var (
	serverOnce sync.Once
	server     *Server
)

func New() *Server {
	e := gin.Default()
	s := Server{engine: e, handler: &Handler{}}
	return &s
}

func GetGin() *Server {
	serverOnce.Do(func() {
		server = New()
	})
	return server
}

func (s *Server) Initialize(address string, user string, pass string) {
	auth := s.engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		user: pass,
		//"user2": "pass2", // user:user2 password:pass2
	}))
	s.authorized = auth
	s.setupRouter()
	go s.run(address)
}

func (s *Server) run(address string) {
	err := s.engine.Run(address)
	if err != nil {
		pkg.GetLog().Fatal(err)
	}
}
