package database

import (
	"fmt"
	"github.com/sathish-30/Golang-fiber/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DBConn *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func Connect(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	} else {
		err := DBConn.AutoMigrate(&model.Book{})
		if err != nil {
			log.Fatal("Could not load database")
		}
		fmt.Println("Database Connected!")
	}
}
