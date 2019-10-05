package grpc

import (
	"google.golang.org/grpc"
	"net"
	api "{{ProjectName}}/api/proto/src"
	"{{ProjectName}}/internal/service"
	"{{ProjectName}}/pkg"
)

type Server struct{}

type Option struct {
	Address string
}

func NewGrpcServer(service *service.ExampleServiceImpl, option Option) *Server {
	go listenGrpc(service, option.Address)
	return &Server{}
}

func listenGrpc(service *service.ExampleServiceImpl, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		pkg.Logger.Fatalf("failed to listen: %v", err)
	}
	pkg.Logger.Info("Start listening on address: ", address)
	s := grpc.NewServer()
	api.RegisterExampleServer(s, NewHandler(service))
	if err := s.Serve(lis); err != nil {
		pkg.Logger.Fatalf("failed to serve: %v", err)
	}
}
