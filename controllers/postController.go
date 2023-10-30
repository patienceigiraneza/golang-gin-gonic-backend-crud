package controllers

import "github.com/gin-gonic/gin"
import "test.com/test/intro/models"
import "test.com/test/intro/initializers"


func PostCreate(c *gin.Context){

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"ID": post.ID,
		"rowsAffected": result.RowsAffected,
		"post": post,

	})

}


func FetchPosts(c *gin.Context){
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts":posts,
	})
}



func GetaPost(c *gin.Context){
	id :=  c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post":post,
	})
}



func UpdateaPost(c *gin.Context){
	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	id :=  c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body : body.Body,

	})

	c.JSON(200, gin.H{
		"post":post,
	})
}


func DeleteaPost(c *gin.Context){
	id :=  c.Param("id")

	var post models.Post
	initializers.DB.Delete(&post, id)
	c.JSON(200, gin.H{
		"status":"deleted",
	})
}
