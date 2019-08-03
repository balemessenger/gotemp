package postgres

import (
	"github.com/jackc/pgx"
	"{{ProjectName}}/internal/repositories"
	"{{ProjectName}}/pkg"
)

type Database struct {
	log          *pkg.Logger
	PostgresPool *pgx.ConnPool
	Example      repositories.ExampleRepo
}

type Option struct {
	Host string
	User string
	Pass string
	Db   string
}

func New(log *pkg.Logger, option Option) *Database {
	var err error
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     option.Host,
			User:     option.User,
			Password: option.Pass,
			Database: option.Db,
		},
		MaxConnections: 10,
	}

	postgresPool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		log.Fatal("Unable to create connection pool", "error", err)
	}
	return &Database{
		log:          log,
		PostgresPool: postgresPool,
		Example:      NewExampleRepo(log, postgresPool),
	}
}

func (d *Database) GetExampleRepo() repositories.ExampleRepo {
	return d.Example
}
