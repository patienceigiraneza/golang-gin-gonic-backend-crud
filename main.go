package main

import "github.com/gin-gonic/gin"
import "test.com/test/intro/initializers"
import "test.com/test/intro/controllers"


func init(){
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.PostCreate)
	r.POST("/", controllers.PostCreate)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
