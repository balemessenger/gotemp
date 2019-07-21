package test

import (
	{{ if Grpc }}
    grpc2 "{{ProjectName}}/api/grpc"
	{{ end }}
	"{{ProjectName}}/api/http"
	"{{ProjectName}}/internal"
	"{{ProjectName}}/pkg"
	"{{ProjectName}}/testkit"
	"math/rand"
	"os"
	"testing"
	"time"
)

func setup() {
	rand.Seed(time.Now().Unix())
	testkit.InitTestConfig("config.yaml")
	pkg.GetLog().Initialize("DEBUG")
	{{ if Grpc }}
	grpc2.GetGrpc().Initialize(internal.GetConfig().Endpoints.Grpc.Address)
	testkit.GetGrpcClient().Initialize(internal.GetConfig().Endpoints.Grpc.Address)
	{{ end }}
	http.GetGin().Initialize(
		internal.GetConfig().Endpoints.Http.Address,
		internal.GetConfig().Endpoints.Http.User,
		internal.GetConfig().Endpoints.Http.Pass)
	time.Sleep(400 * time.Millisecond)
}

func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	r := m.Run()
	teardown()
	os.Exit(r)
}