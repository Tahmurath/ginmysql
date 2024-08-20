package services

import (
	"resedist/models"
	"resedist/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{
		userRepo: repositories.NewUserRepository(),
	}
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
	return us.userRepo.FindByID(id)
}
