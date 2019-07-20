package cmd

import (
	"fmt"
	"{{ProjectName}}/api/grpc"
	"{{ProjectName}}/api/http"
	"{{ProjectName}}/pkg"
	"{{ProjectName}}/pkg/postgres"
)

func initialize() {
	fmt.Println("{{ProjectName}} build version:", pkg.BuildVersion)
	fmt.Println("{{ProjectName}} build time:", pkg.BuildTime)
	pkg.GetConfig().Initialize("")
	pkg.GetLog().Initialize(pkg.GetConfig().Conf.Log.Level)
	postgres.GetPostgresDB()
	grpc.Initialize()
	http.GetGin().Initialize(pkg.GetConfig().Conf.Endpoints.Http.Address)
	pkg.GetPrometheus().Initialize(pkg.GetConfig().Conf.Prometheus.Port)
}

func Main() {
	initialize()
	// Write your entryoint here
	pkg.Signal.Wait()
}
