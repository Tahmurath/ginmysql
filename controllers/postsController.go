package controllers

import (
	"net/http"
	"vida/initializers"
	"vida/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var posted models.Post
	//c.Bind(&posted)

	if err := c.ShouldBindJSON(&posted); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var posted models.Post
	//c.Bind(&posted)

	if err := c.ShouldBindJSON(&posted); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: posted.Title,
		Body:  posted.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	//var post models.Post
	//initializers.DB.First(&post, id)

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
