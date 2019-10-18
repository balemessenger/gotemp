package test

import (
	grpc2 "{{ProjectName}}/api/grpc"
	"{{ProjectName}}/internal"

	"{{ProjectName}}/api/http"
	"{{ProjectName}}/internal/service"
	"{{ProjectName}}/pkg"
	"{{ProjectName}}/testkit"
	"math/rand"
	"os"
	"testing"
	"time"
)

var Conf *internal.Config

func setup() {
	rand.Seed(time.Now().Unix())
	Conf = testkit.InitTestConfig("config.yaml")
	pkg.Logger.SetLevel(Conf.Log.Level)

	srv := service.NewExampleService() // Inject dependencies here
	grpc2.NewGrpcServer(srv, grpc2.Option{
		Address: Conf.Endpoints.Grpc.Address,
	})

	testkit.GetGrpcClient().Initialize(Conf.Endpoints.Grpc.Address)

	http.NewHttpServer(
		http.Option{
			Address: Conf.Endpoints.Http.Address,
			User:    Conf.Endpoints.Http.User,
			Pass:    Conf.Endpoints.Http.Pass,
		})

	time.Sleep(4000 * time.Millisecond)
}

func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	r := m.Run()
	teardown()
	os.Exit(r)
}
