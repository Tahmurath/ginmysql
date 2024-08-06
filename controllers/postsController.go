package controllers

import (
	"vida/initializers"
	"vida/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// var body struct {
	// 	Body  string
	// 	Title string
	// }
	var posted models.Post
	c.Bind(&posted)

	post := models.Post{Title: posted.Title, Body: posted.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}
