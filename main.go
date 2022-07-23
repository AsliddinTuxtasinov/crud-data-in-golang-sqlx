package main

import (
	"go-with-db/controllers"
	"go-with-db/db_client"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-with-db/docs"
)

// @title           Simple CRUD app in golang
// @version         1.0
// @description     This is a simple CRUD app in golang for learning.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://google.com
// @contact.email  asliddintukhtasinov5@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used
func main() {

	db_client.InitialiseDBConnection()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		post := v1.Group("/post")
		{
			post.POST("/", controllers.CreatePost)
			post.GET("/", controllers.GetPosts)
			post.GET("/:id", controllers.GetPost)
			post.DELETE("/:id", controllers.DeletePost)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
