package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-api/models"
	"go-rest-api/repository"
	"net/http"
	"strconv"
)

var UserRepo repository.UserRepository

func CreateUser(c *gin.Context) {
	var input models.User
	fmt.Println("Creating a new user")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email}

	if err := UserRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users, err = UserRepo.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {

	userIdInt, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println("Error converting user id to int", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := UserRepo.GetUserById(userIdInt)

	if err != nil {
		fmt.Println("Error getting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {

	var userIdInt, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error converting user id to int", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// gets the current user using the userId
	user, err := UserRepo.GetUserById(userIdInt)

	// returns an error if the user is not found.
	if err != nil {
		fmt.Println("Error getting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to find user!"})
	}

	// Bind the incoming JSON to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Error binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//update the user with the incoming user data
	if err := UserRepo.UpdateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to update user!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var userIdInt, err = strconv.Atoi(c.Param("id"))
	fmt.Printf("Attempting to delete user with id %v\n", userIdInt)

	if err != nil {
		fmt.Println("Error converting user id to int", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// gets the current user using the userId
	user, err := UserRepo.GetUserById(userIdInt)

	// returns an error if the user is not found.
	if err != nil {
		fmt.Println("Error getting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to find user!"})
	}

	if err := UserRepo.DeleteUser(user); err != nil {
		fmt.Println("Error deleting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to delete user!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User with id '" + strconv.Itoa(userIdInt) + "' deleted"})
}
