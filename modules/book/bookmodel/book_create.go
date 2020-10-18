package bookmodel

import (
	"gorm.io/gorm"
	"time"
)

type BookCreate struct {
	gorm.Model
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
	ImageUrl      string    `json:"image_url"`
	Description   string    `json:"description"`
}

func (BookCreate) TableName() string {
	return "books"
}
