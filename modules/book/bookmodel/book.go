package bookmodel

import (
	"versla/sdk/sdkcm"
)

type Book struct {
	sdkcm.Model
	Title  string `json:"title" gorm:"column:title"`
	Author string `json:"author" gorm:"column:author"`
}

func (Book) TableName() string {
	return "books"
}
