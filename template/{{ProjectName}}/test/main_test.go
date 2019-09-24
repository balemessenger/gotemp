package test

import (
	"math/rand"
	"os"
	"testing"
	"time"
	grpc2 "{{ProjectName}}/api/grpc"
	"{{ProjectName}}/api/http"
	"{{ProjectName}}/internal"
	"{{ProjectName}}/pkg"
	"{{ProjectName}}/testkit"
)

var Conf *internal.Config

func setup() {
	rand.Seed(time.Now().Unix())
	Conf = testkit.InitTestConfig("config.yaml")
	log := pkg.NewLog("DEBUG")

	grpc2.New(log, grpc2.Option{
		Address: Conf.Endpoints.Grpc.Address,
	})

	testkit.GetGrpcClient().Initialize(Conf.Endpoints.Grpc.Address)

	http.New(
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
