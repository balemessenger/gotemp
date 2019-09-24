package repositories

import (
	"{{ProjectName}}/pkg/db"
)

type ExampleRepo interface {
	GetAllExampleIds() []int32
}

type ExampleRepoImpl struct {
	db *db.PostgresDb
}

func NewExampleRepo(db *db.PostgresDb) *ExampleRepoImpl {
	return &ExampleRepoImpl{db: db}
}

func (p *ExampleRepoImpl) GetAllExampleIds() []int32 {
	return nil
}
