package postgres

import (
	"github.com/jackc/pgx"
	"{{ProjectName}}/pkg"
)

type ExamplePostgresRepo struct {
	log          *pkg.Logger
	PostgresPool *pgx.ConnPool
}

func NewExampleRepo(log *pkg.Logger, PostgresPool *pgx.ConnPool) *ExamplePostgresRepo {
	return &ExamplePostgresRepo{log, PostgresPool}
}

func (p *ExamplePostgresRepo) GetAllExampleIds() []int32 {
	rows, err := p.PostgresPool.Query("SELECT id FROM users WHERE is_bot=false")
	var users []int32
	switch err {
	case nil:
	case pgx.ErrNoRows:
		return users
	default:
		p.log.Error("Error in GetAllExamples", err)
		return users
	}

	if rows == nil {
		p.log.Error("Error in GetAllExamples: rows is empty")
		return users
	}

	for rows.Next() {
		var u int32
		if e := rows.Scan(&u); e != nil {
			p.log.Error(e)
		}
		users = append(users, u)
	}
	return users
}
