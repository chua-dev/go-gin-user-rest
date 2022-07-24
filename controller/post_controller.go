package controller

import (
	"net/http"

	"github.com/chua-dev/go-gin-user-rest/database"
	"github.com/gin-gonic/gin"
)

// Can be understand as Service

type Post struct {
	Id      int64  `json:"UserId"`
	Topic   string `json:"topic"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var postList = []Post{}

// Get All Posts
func getPosts(c *gin.Context) {
	var postList []Post

	// DBClient.Query aim to query multiple rows
	rows, err := database.DBClient.Query("SELECT * FROM posts;")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid Request Body",
		})
		return
	}

	for rows.Next() {
		var eachPost Post

		// Scan each row into eachPost struct
		if err := rows.Scan(&eachPost.Id, &eachPost.Topic, &eachPost.Title, &eachPost.Content); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   true,
				"message": "Database Error",
			})
			return
		}

		postList = append(postList, eachPost)
	}

	c.JSON(http.StatusOK, postList)

}

// Create User

// Get Post By Id
