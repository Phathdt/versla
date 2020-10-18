package bookmodel

import (
	"time"
)

type BookUpdate struct {
	Title         *string    `json:"title" gorm:"column:title"`
	Author        *string    `json:"author" gorm:"column:author"`
	PublishedDate *time.Time `json:"published_date" gorm:"column:published_date"`
	ImageUrl      *string    `json:"image_url" gorm:"column:image_url"`
	Description   *string    `json:"description" gorm:"column:description"`
}

func (BookUpdate) TableName() string {
	return "books"
}
