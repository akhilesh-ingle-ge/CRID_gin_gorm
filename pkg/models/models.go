package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var DB *gorm.DB

var dsn = "postgres://postgres:12345@localhost:5432/test?sslmode=disable"

func ConnectToDB() {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&Book{})
	fmt.Println("Connection Successful")

	DB = db
}
