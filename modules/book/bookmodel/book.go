package bookmodel

import (
	"time"
	"versla/util/gorm"
)

type Book struct {
	gorm.Model
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
	ImageUrl      string    `json:"image_url"`
	Description   string    `json:"description"`
}

func (Book) TableName() string {
	return "books"
}
