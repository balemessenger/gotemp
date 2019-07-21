package postgres

import (
	"{{ProjectName}}/pkg"
	"os"
	"sync"

	"github.com/jackc/pgx"
)

type Database struct {
	PostgresPool *pgx.ConnPool
}

var (
	once             sync.Once
	postgresDatabase *Database
)

func GetPostgres() *Database {
	once.Do(func() {
		postgresDatabase = New()
	})
	return postgresDatabase
}

func New() *Database {
	return &Database{}
}

func (p *Database) Initialize(host string, user string, pass string, db string) *pgx.ConnPool {
	var err error
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     host,
			User:     user,
			Password: pass,
			Database: db,
		},
		MaxConnections: 10,
	}

	postgresPool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		pkg.GetLog().Error("Unable to create connection pool", "error", err)
		os.Exit(1)
	}
	p.PostgresPool = postgresPool
	return postgresPool
}
