package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=docker.local user=go_user password=password dbname=go_rest_api port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database!")
	}

	DB = database
	fmt.Println("Successfully connected to the database!")

}
