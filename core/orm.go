package core

import (
	"fmt"
	"goginapi/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Orm struct {
}

func (core Orm) Done() *gorm.DB {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to find .env file")
	}

	// migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Book{},
	)
}
