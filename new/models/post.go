package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `json:"post_title" binding:"required"`
	Body  string `json:"post_body"`
}
