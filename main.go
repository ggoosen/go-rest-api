package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-api/config"
	"go-rest-api/handlers"
	"go-rest-api/middleware"
	"go-rest-api/models"
	"go-rest-api/repository"
)

func main() {

	config.ConnectDatabase()
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		fmt.Println(err)
		return
	}

	userRepo := repository.NewGormUserRepository(config.DB)
	handlers.UserRepo = userRepo

	r := gin.Default()

	// Apply the validation error handler middleware globally
	r.Use(middleware.ValidationErrorHandler())

	r.POST("/users", handlers.CreateUser)
	r.GET("/user/:id", handlers.GetUser)
	r.GET("/users", handlers.GetUsers)
	r.DELETE("/user/:id", handlers.DeleteUser)
	r.PUT("/user/:id", handlers.UpdateUser)

	if err := r.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
