package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint `gorm:"primaryKey"`
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	Todo struct {
		ID     uint `gorm:"primaryKey"`
		Text   string
		Done   bool
		UserID uint
		User   User
	}
)

func Setup() (*gorm.DB, error) {
	dsn := "root:sample@tcp(example-go-graphql-mysql:3306)/graphql_go_example?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
