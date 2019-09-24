package repositories

import (
	"github.com/jinzhu/gorm"
	"{{ProjectName}}/pkg"
)

type Database struct {
	Log *pkg.Logger
	Db  *gorm.DB
}

type DatabaseRepo interface {
	GetExampleRepo() ExampleRepo
}
