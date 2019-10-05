package cmd

import (
	"fmt"
	"time"
	"{{ProjectName}}/api/grpc"
	"{{ProjectName}}/api/http"
	"{{ProjectName}}/internal"
	"{{ProjectName}}/internal/service"
	"{{ProjectName}}/pkg"
	"{{ProjectName}}/pkg/db"
	{{ if Kafka }}
	"{{ProjectName}}/pkg/pubsub"
	{{ end }}
)

type Server struct {
	isReady chan bool
}

func NewServer() *Server {
	fmt.Println("{{ProjectName}} build version:", pkg.BuildVersion)
	fmt.Println("{{ProjectName}} build time:", pkg.BuildTime)
	conf := internal.NewConfig("")
	pkg.NewLog(conf.Log.Level)

	{{ if Postgres }}
	_ = db.NewPostgres(db.PostgresConfig{
		Host: conf.Postgres.Host,
		Port: conf.Postgres.Port,
		User: conf.Postgres.User,
		Pass: conf.Postgres.Pass,
		Db:   conf.Postgres.DB,
	})
	{{ end }}
	{{ if Cassandra }}
	_ = db.NewCassandra(db.CassConfig{
		Hosts:         conf.Cassandra.Hosts,
		Port:          conf.Cassandra.Port,
		Password:      conf.Cassandra.Password,
		Username:      conf.Cassandra.Username,
		KeySpace:      conf.Cassandra.KeySpace,
		Consistency:   conf.Cassandra.Consistency,
		PageSize:      conf.Cassandra.PageSize,
		Timeout:       time.Duration(conf.Cassandra.Timeout) * time.Millisecond,
		DataCenter:    conf.Cassandra.DataCenter,
		PartitionSize: conf.Cassandra.PartitionSize,
	})
	{{ end }}
	{{ if Kafka }}
	_ = pubsub.NewKafka(
		pubsub.KafkaOption{
			Servers:     conf.Kafka.BootstrapServers,
			GroupId:     conf.Kafka.GroupId,
			OffsetReset: conf.Kafka.AutoOffsetReset,
		})
	{{ end }}
	// TODO: Init repositories here
	srv := service.NewExampleService() // Inject dependencies here
	{{ if Grpc }}
	grpc.NewGrpcServer(srv, grpc.Option{
		Address: conf.Endpoints.Grpc.Address,
	})
	{{ end }}

	http.NewHttpServer(
		http.Option{
			Address: conf.Endpoints.Http.Address,
			User:    conf.Endpoints.Http.User,
			Pass:    conf.Endpoints.Http.Pass,
		})

	pkg.NewPrometheus(conf.Prometheus.Port)

	return &Server{
		isReady: make(chan bool),
	}
}

func (s *Server) Run() bool {
	go s.start()
	return <-s.isReady
}

func (s *Server) start() {
	// Write your entry point
	s.isReady <- true
}

func Main() {
	server := NewServer()
	pkg.Logger.Info("Hello {{ProjectName}}")
	server.Run()
	pkg.Signal.Wait()
}
