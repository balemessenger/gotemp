package postgres

import (
	"github.com/jackc/pgx"
	"{{ProjectName}}/pkg"
)

type ExamplePostgresRepo struct {
	PostgresPool *pgx.ConnPool
}

func NewExampleRepo(PostgresPool *pgx.ConnPool) *ExamplePostgresRepo {
	return &ExamplePostgresRepo{PostgresPool}
}

func (u *ExamplePostgresRepo) GetAllExampleIds() []int32 {
	rows, err := u.PostgresPool.Query("SELECT id FROM users WHERE is_bot=false")
	var users []int32
	switch err {
	case nil:
	case pgx.ErrNoRows:
		return users
	default:
		pkg.GetLog().Error("Error in GetAllExamples", err)
		return users
	}

	if rows == nil {
		pkg.GetLog().Error("Error in GetAllExamples: rows is empty")
		return users
	}

	for rows.Next() {
		var u int32
		if e := rows.Scan(&u); e != nil {
			pkg.GetLog().Error(e)
		}
		users = append(users, u)
	}
	return users
}
