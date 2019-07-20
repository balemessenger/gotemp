package grpc

import (
	//"google.golang.org/grpc"
	//api "{{ProjectName}}/api/proto/src"
	//"log"
	//"net"
	"{{ProjectName}}/pkg"
)

type Server struct{}

func Initialize() {
	go ListenGrpc(pkg.GetConfig().Conf.Endpoints.Grpc.Address)
}

func ListenGrpc(address string) {
	//lis, err := net.Listen("tcp", address)
	//if err != nil {
	//	log.Logger.Fatalf("failed to listen: %v", err)
	//}
	//log.Logger.Info("Start listening on address: ", address)
	//s := grpc.NewServer()
	//api.RegisterPushServer(s, &Server{})
	//if err := s.Serve(lis); err != nil {
	//	log.Logger.Fatalf("failed to serve: %v", err)
	//}
}
