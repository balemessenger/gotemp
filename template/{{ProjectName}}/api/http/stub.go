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

func New(log *pkg.Logger, option Option) *Server {
	engine := gin.Default()
	auth := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		option.User: option.Pass,
		//"user2": "pass2", // user:user2 password:pass2
	}))
	s := Server{
		engine:     engine,
		authorized: auth,
		handler:    NewHandler(log)}
	s.setupRouter()

	go func(address string) {
		err := s.engine.Run(address)
		if err != nil {
			log.Fatal(err)
		}
	}(option.Address)

	return &s
}
