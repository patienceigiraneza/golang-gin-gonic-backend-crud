package main

import "github.com/gin-gonic/gin"
import "test.com/test/intro/initializers"
import "test.com/test/intro/controllers"


func init(){
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.FetchPosts)
	r.POST("/", controllers.PostCreate)
	r.GET("/:id", controllers.GetaPost)
	r.PUT("/:id", controllers.UpdateaPost)
	r.DELETE("/:id", controllers.DeleteaPost)

	r.Run(":8000")
}
