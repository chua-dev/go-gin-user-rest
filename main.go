package main

import (
	"log"
	"net/http"

	"github.com/chua-dev/go-gin-user-rest/controller"
	"github.com/chua-dev/go-gin-user-rest/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// User Model Mapping Structure
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Initialize fake database
var userList []User      // Return Null
var userList2 = []User{} // Return Empty Array []

func main() {
	database.ConnectDatabase()

	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controller.GetUsers)
		userRoutes.GET("/:id", controller.GetUserById)
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.PUT("/:id", EditUser) // /users/123
		userRoutes.DELETE("/:id", DeleteUser)
	}

	//router.GET("/", GetUsers)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err.Error())
	}
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

func EditUser(c *gin.Context) {
	id := c.Param("id") // Get params as string
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for index, user := range userList {
		if user.ID == id {
			userList[index].Name = reqBody.Name
			userList[index].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for index, user := range userList {
		if user.ID == id {
			// Get all user before i
			userList = append(userList[:index], userList[index+1:]...)
			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
