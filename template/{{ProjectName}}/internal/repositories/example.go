package repositories

import "github.com/jinzhu/gorm"

type Example struct {
	gorm.Model
	UserId   int `gorm:"primary_key"`
	Username string
	Password string
}

type ExampleRepo interface {
	GetAllExampleIds() []int32
}
