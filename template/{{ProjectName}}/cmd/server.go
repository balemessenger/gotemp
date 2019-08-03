package cmd

import (
	"fmt"
	"{{ProjectName}}/api/grpc"

	"{{ProjectName}}/api/http"
	"{{ProjectName}}/internal"
	"{{ProjectName}}/internal/server"
	"{{ProjectName}}/pkg"

	"{{ProjectName}}/internal/postgres"

	"{{ProjectName}}/pkg/pubsub"
)

func initialize() *pkg.Logger {
	fmt.Println("{{ProjectName}} build version:", pkg.BuildVersion)
	fmt.Println("{{ProjectName}} build time:", pkg.BuildTime)
	conf := internal.NewConfig("")
	log := pkg.NewLog(conf.Log.Level)

	db := postgres.New(log, postgres.Option{
		Host: conf.Postgres.Host,
		User: conf.Postgres.User,
		Pass: conf.Postgres.Pass,
		Db:   conf.Postgres.DB,
	})

	kafka := pubsub.NewKafka(
		log,
		pubsub.KafkaOption{
			Servers:     conf.Kafka.BootstrapServers,
			GroupId:     conf.Kafka.GroupId,
			OffsetReset: conf.Kafka.AutoOffsetReset,
		})

	grpc.New(log, grpc.Option{
		Address: conf.Endpoints.Grpc.Address,
	})

	http.New(
		log,
		http.Option{
			Address: conf.Endpoints.Http.Address,
			User:    conf.Endpoints.Http.User,
			Pass:    conf.Endpoints.Http.Pass,
		})

	pkg.NewPrometheus(log, conf.Prometheus.Port)

	//Initialize main logic
	internal.NewExample(log, db, kafka).Start(conf.Core.WorkPoolSize)

	return log
}

func Main() {
	log := initialize()
	log.Info("Hello {{ProjectName}}")
	server.New().Run()
	pkg.Signal.Wait()
}
