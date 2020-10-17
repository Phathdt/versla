package bookstorage

import "gorm.io/gorm"

type bookSQLStorage struct {
	DB *gorm.DB
}

func NewBookSQLStorage(DB *gorm.DB) *bookSQLStorage {
	return &bookSQLStorage{
		DB: DB,
	}
}
