package repositories

import (
	"errors"
	"resedist/models"
)

type PostRepository struct {
	// اینجا می‌توانید دیتابیس را به عنوان وابستگی اضافه کنید
}

func NewPostRepository() PostRepository {
	return PostRepository{}
}

func (ur *PostRepository) FindByID(id string) (*models.Post, error) {
	// اینجا می‌توانید کوئری واقعی دیتابیس را بنویسید.
	// برای مثال، اینجا یک کاربر فرضی برگردانده می‌شود.

	if id == "1" {
		return &models.Post{
			Title: "John Doe",
			Body:  "johndoe@example.com",
		}, nil
	}
	return nil, errors.New("post not found")
}
