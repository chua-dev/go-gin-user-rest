package controller

import (
	"net/http"

	"github.com/chua-dev/go-gin-user-rest/database"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func GetUsers(c *gin.Context) {
	var users []User

	// Query Multiple Row
	rows, err := database.DBClient.Query("SELECT ID, Name, Age, Email from users;")

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	for rows.Next() {
		var singleUser User
		// Scan is Map the row data
		if err := rows.Scan(&singleUser.ID, &singleUser.Name, &singleUser.Age, &singleUser.Email); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": true,
			})
			return
		}

		users = append(users, singleUser)
	}

	c.JSON(http.StatusOK, users)

}

func GetUserById(c *gin.Context) {
	var selectedUser User
	id := c.Param("id")
	row := database.DBClient.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if err := row.Scan(&selectedUser.ID, &selectedUser.Name, &selectedUser.Age, &selectedUser.Email); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, selectedUser)
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

	// Execute SQL query
	// ? scape value, prevent SQL injection Like DELETE into VALUE
	res, err := database.DBClient.Exec("INSERT INTO users(Name, Age, Email) VALUES (?, ?, ?);",
		reqBody.Name,
		reqBody.Age,
		reqBody.Email,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return
	}

	id, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()

	c.JSON(http.StatusCreated, gin.H{
		"status":                 "success",
		"id":                     id,
		"number of row affected": rowsAffected,
	})
}
