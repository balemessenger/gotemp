package postgres

import (
	"{{ProjectName}}/pkg"
	"os"
	"sync"

	"github.com/jackc/pgx"
)

type PostgresDatabase struct {
	PostgresPool *pgx.ConnPool
}

var (
	once             sync.Once
	postgresDatabase *PostgresDatabase
)

func GetPostgresDB() *PostgresDatabase {
	once.Do(func() {
		postgresDatabase = New()
	})
	return postgresDatabase
}

func New() *PostgresDatabase {
	return &PostgresDatabase{}
}

func (p *PostgresDatabase) Initialize() *pgx.ConnPool {
	var err error
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     pkg.GetConfig().Conf.Postgres.Host,
			User:     pkg.GetConfig().Conf.Postgres.User,
			Password: pkg.GetConfig().Conf.Postgres.Pass,
			Database: pkg.GetConfig().Conf.Postgres.DB,
		},
		MaxConnections: 10,
	}

	postgresPool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		pkg.GetLog().Logger.Error("Unable to create connection pool", "error", err)
		os.Exit(1)
	}
	p.PostgresPool = postgresPool
	return postgresPool
}
