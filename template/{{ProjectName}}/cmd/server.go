package cmd

import (
	"fmt"
	{{ if Grpc }}
	"{{ProjectName}}/api/grpc"
	{{ end }}
	"{{ProjectName}}/api/http"
	"{{ProjectName}}/internal"
	"{{ProjectName}}/internal/server"
	"{{ProjectName}}/pkg"
	{{ if Postgres }}
	"{{ProjectName}}/pkg/postgres"
	{{ end }}
    {{ if Kafka }}
	"{{ProjectName}}/pkg/pubsub"
	{{ end }}
)

func initialize() {
	fmt.Println("{{ProjectName}} build version:", pkg.BuildVersion)
	fmt.Println("{{ProjectName}} build time:", pkg.BuildTime)
	internal.GetConfig().Initialize("")
	pkg.GetLog().Initialize(internal.GetConfig().Log.Level)
	{{ if Postgres }}
	postgres.GetPostgres().Initialize(
		internal.GetConfig().Postgres.Host,
		internal.GetConfig().Postgres.User,
		internal.GetConfig().Postgres.Pass,
		internal.GetConfig().Postgres.DB)
	{{ end }}
	{{ if Kafka }}
	pubsub.GetKafka().Initialize(
		internal.GetConfig().Kafka.BootstrapServers,
		internal.GetConfig().Kafka.GroupId,
		internal.GetConfig().Kafka.AutoOffsetReset)
	{{ end }}
	{{ if Grpc }}
	grpc.GetGrpc().Initialize(internal.GetConfig().Endpoints.Grpc.Address)
	{{ end }}
	http.GetGin().Initialize(
		internal.GetConfig().Endpoints.Http.Address,
		internal.GetConfig().Endpoints.Http.User,
		internal.GetConfig().Endpoints.Http.Pass)
	pkg.GetPrometheus().Initialize(internal.GetConfig().Prometheus.Port)

	//Initialize main logic
	internal.GetExample().Initialize(internal.GetConfig().Core.WorkPoolSize)
}

func Main() {
	initialize()
	pkg.GetLog().Info("Hello {{ProjectName}}")
	server.New().Run()
	pkg.Signal.Wait()
}
