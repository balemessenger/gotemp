package testkit

import (
	"google.golang.org/grpc"
	api "{{ProjectName}}/api/proto/src"
	"log"
	"sync"
)

var (
	grpcOnce sync.Once
	grpcClient  *GrpcClient
)

type GrpcClient struct {
	api.ExampleClient
}

func NewGrpcClient() *GrpcClient {

	return &GrpcClient{}
}

func GetGrpcClient() *GrpcClient{
	grpcOnce.Do(func() {
		grpcClient = NewGrpcClient()
	})
	return grpcClient
}

func (c *GrpcClient) Initialize(address string){
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	client := api.NewExampleClient(conn)
	c.ExampleClient = client
}
