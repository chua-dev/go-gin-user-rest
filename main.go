package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Initialize fake database
var userList []User      // Return Null
var userList2 = []User{} // Return Empty Array []

func main() {
	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
	}

	//router.GET("/", GetUsers)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err.Error())
	}
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

func CreateUser(c *gin.Context) {
	var reqBody User
	// Pointer of a struct object as param
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	userList = append(userList, reqBody)
	c.JSON(200, gin.H{
		"error": false,
	})
}
