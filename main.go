package main

import (
	"go-with-db/controllers"
	"go-with-db/db_client"

	"github.com/gin-gonic/gin"
)

func main() {
	db_client.InitialiseDBConnection()

	r := gin.Default()
	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.DELETE("/:id", controllers.DeletePost)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
