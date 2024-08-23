package services

import (
	"resedist/models"
	"resedist/repositories"
)

type PostService struct {
	postRepo repositories.PostRepository
}

func NewPostService() PostService {
	return PostService{
		postRepo: repositories.NewPostRepository(),
	}
}

func (us *PostService) GetPostByID(id string) (*models.Post, error) {
	return us.postRepo.FindByID(id)
}
