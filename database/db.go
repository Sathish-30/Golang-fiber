package database

import (
	"fmt"
	"github.com/sathish-30/Golang-fiber/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("handler.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("Database Connection successfully opened")

	DBConn.AutoMigrate(&model.Book{})
	fmt.Println("DataBase Migrated")
}
