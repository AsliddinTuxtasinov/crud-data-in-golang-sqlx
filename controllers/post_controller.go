package controllers

import (
	"fmt"
	"go-with-db/db_client"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int64     `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func CreatePost(c *gin.Context) {
	var reqBody Post

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid request body",
		})
		log.Println(err.Error())
		return
	}

	res, err := db_client.DBClient.Exec(
		"INSERT INTO posts(title, content) VALUES (?, ?);",
		reqBody.Title, reqBody.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Status Internal Server Error",
		})
		log.Println(err.Error())
		return
	}

	id, _ := res.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Status Created",
		"post": gin.H{
			"id":         id,
			"title":      reqBody.Title,
			"content":    reqBody.Content,
			"created_at": reqBody.CreatedAt,
		},
	})
	return
}

func GetPosts(c *gin.Context) {
	var posts []Post

	err := db_client.DBClient.Select(&posts, "SELECT * FROM posts;") // Slice of Rows (queryRow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Status Internal Server Error",
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
	return
}

func GetPost(c *gin.Context) {
	var singlePost Post
	id := c.Param("id")

	err := db_client.DBClient.Get(&singlePost, "SELECT * FROM posts WHERE id=?;", id) // Single Row (queryRow)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": singlePost,
	})
	return
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	res, err := db_client.DBClient.Exec("DELETE FROM posts WHERE id=?;", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": fmt.Sprintf("Deleted post %v", count),
	})
	return

}
