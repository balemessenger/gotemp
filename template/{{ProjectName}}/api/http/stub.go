package http

import (
	"github.com/gin-gonic/gin"
	"{{ProjectName}}/pkg"
)

type Server struct {
	engine     *gin.Engine
	authorized *gin.RouterGroup
	handler    *Handler
}

type Option struct {
	Address string
	User    string
	Pass    string
}

func NewHttpServer(option Option) *Server {
	engine := gin.Default()
	auth := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		option.User: option.Pass,
	}))
	s := Server{
		engine:     engine,
		authorized: auth,
		handler:    NewHttpHandler(),
	}
	s.setupRouter()

	go func(address string) {
		err := s.engine.Run(address)
		if err != nil {
			pkg.Logger.Fatal(err)
		}
	}(option.Address)

	return &s
}
