package postgres

import (
	"{{ProjectName}}/pkg"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

type Example struct {
	gorm.Model
	UserId   int `gorm:"primary_key"`
	Username string
	Password string
}

type Database struct {
	log *pkg.Logger
	db  *gorm.DB
}

type Option struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

func New(log *pkg.Logger, option Option) *Database {
	url := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", option.Host, option.Port, option.User, option.Db, option.Pass)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&Example{})

	return &Database{
		log: log,
		db:  db,
	}
}
