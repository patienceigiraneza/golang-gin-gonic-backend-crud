package main

import (
    "github.com/gin-gonic/gin"
)

var names = map[string]string{
    "fname": "Patience",
    "lname": "Igiraneza",
}



func getName(c *gin.Context) {

    firstName := names["fname"]
    lastName := names["lname"]

    c.JSON(200, gin.H{
        "first_name": firstName,
        "last_name":  lastName,
    })
}


func createName(c *gin.Context){
	first_name := c.PostForm("fname")
	last_name := c.PostForm("lname")

	if (len(first_name) == 0 || len(last_name) == 0){
		c.JSON(404, gin.H{
			"fname": first_name,
			"lname": last_name,
			"status": "Post",
		})
		return
	}

}

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "golang is awsome",
        })
    })

    r.GET("/name", getName)
    r.POST("/name", createName)

    r.Run(":8000")
}
