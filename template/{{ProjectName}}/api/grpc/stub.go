package grpc

import (
	"google.golang.org/grpc"
	api "{{ProjectName}}/api/proto/src"
	"{{ProjectName}}/pkg"
	"net"
	"sync"
)

type Server struct{}

var (
	serverOnce sync.Once
	server     *Server
)

func New() *Server {
	return &Server{}
}

func GetGrpc() *Server {
	serverOnce.Do(func() {
		server = New()
	})
	return server
}

func (s *Server) Initialize(address string) {
	go s.listenGrpc(address)

}

func (*Server) listenGrpc(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		pkg.GetLog().Fatalf("failed to listen: %v", err)
	}
	pkg.GetLog().Info("Start listening on address: ", address)
	s := grpc.NewServer()
	api.RegisterExampleServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		pkg.GetLog().Fatalf("failed to serve: %v", err)
	}
}
