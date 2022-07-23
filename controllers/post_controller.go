package controllers

import (
	"fmt"
	"go-with-db/db_client"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PostCreate struct {
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}

type Post struct {
	PostCreate
	ID        int64     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// AddPost godoc
// @Summary      Add a post
// @Description  add by json post
// @Tags         post
// @Accept       json
// @Produce      json
// @Param        post  body PostCreate  true  "Add account"
// @Success      200   {object}  Post
// @Router       /post [post]
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
			"created_at": time.Now(),
		},
	})
	return
}

// ListPostss godoc
// @Summary      List posts
// @Description  get posts
// @Tags         post
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(title)
// @Success      200  {object}  Post
// @Router       /post [get]
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

// ShowPost godoc
// @Summary      Show an post
// @Description  get post by ID
// @Tags         post
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  Post
// @Router       /post/{id} [get]
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

// DeletePost godoc
// @Summary      Delete post
// @Description  Delete post by ID
// @Tags         post
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "post ID"  Format(int64)
// @Success      204  {string}  http.StatusOK
// @Router       /post/{id} [delete]
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
