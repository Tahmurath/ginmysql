package repositories

import (
	"errors"
	"resedist/models"
)

type UserRepository struct {
	// اینجا می‌توانید دیتابیس را به عنوان وابستگی اضافه کنید
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ur *UserRepository) FindByID(id string) (*models.User, error) {
	// اینجا می‌توانید کوئری واقعی دیتابیس را بنویسید.
	// برای مثال، اینجا یک کاربر فرضی برگردانده می‌شود.

	if id == "1" {
		return &models.User{
			ID:    "1",
			Name:  "John Doe",
			Email: "johndoe@example.com",
		}, nil
	}
	return nil, errors.New("user not found")
}
