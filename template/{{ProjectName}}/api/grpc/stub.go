package grpc

import (
	"google.golang.org/grpc"
	api "{{ProjectName}}/api/proto/src"
	"{{ProjectName}}/pkg"
	"net"
)

type Server struct{}

type Option struct {
	Address string
}

func New(log *pkg.Logger, option Option) *Server {
	go listenGrpc(log, option.Address)
	return &Server{}
}

func listenGrpc(log *pkg.Logger, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Info("Start listening on address: ", address)
	s := grpc.NewServer()
	api.RegisterExampleServer(s, NewHandler(log))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
