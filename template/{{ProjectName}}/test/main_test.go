package test

import (
	grpc2 "{{ProjectName}}/api/grpc"
	"{{ProjectName}}/internal"

	"{{ProjectName}}/api/http"
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
	log := pkg.NewLog("DEBUG")

	grpc2.NewGrpcServer(log, grpc2.Option{
		Address: Conf.Endpoints.Grpc.Address,
	})

	testkit.GetGrpcClient().Initialize(Conf.Endpoints.Grpc.Address)

	http.NewHttpServer(
		log,
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
