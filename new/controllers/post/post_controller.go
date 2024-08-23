package post

import (
	"net/http"
	"resedist/services"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (uc *PostController) GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := uc.postService.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}
