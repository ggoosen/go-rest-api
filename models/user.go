package models

import (
	config "go-rest-api/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" gorm:"unique" binding:"required,email"`
}

func init() {
	config.ConnectDatabase()
}
