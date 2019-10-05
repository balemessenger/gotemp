package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"{{ProjectName}}/pkg"
)

type PostgresDb struct {
	db *gorm.DB
}

type PostgresConfig struct {
	Host string
	Port int
	User string
	Pass string
	Db   string
}

func NewPostgres(PosgresConfig PostgresConfig) *PostgresDb {
	url := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", PosgresConfig.Host, PosgresConfig.Port, PosgresConfig.User, PosgresConfig.Db, PosgresConfig.Pass)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		pkg.Logger.Fatal("failed to connect database", err)
	}
	db.LogMode(true)

	return &PostgresDb{
		db: db,
	}
}
