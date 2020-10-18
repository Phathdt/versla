package bookmodel

import (
	"versla/sdk/sdkcm"
)

type BookCreate struct {
	sdkcm.Model
	Title  string `json:"title" gorm:"column:title"`
	Author string `json:"author" gorm:"column:author"`
}

func (BookCreate) TableName() string {
	return "books"
}
