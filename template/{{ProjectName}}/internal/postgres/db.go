package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"{{ProjectName}}/internal/repositories"
	"{{ProjectName}}/pkg"
)

type Option struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

func New(log *pkg.Logger, option Option) *repositories.Database {
	url := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", option.Host, option.Port, option.User, option.Db, option.Pass)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&repositories.Example{})

	return &repositories.Database{
		Log: log,
		Db:  db,
	}
}
