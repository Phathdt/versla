package bookstorage

import (
	"context"
	"versla/modules/book/bookmodel"
)

func (storage bookSQLStorage) CreateBook(ctx context.Context, data *bookmodel.BookCreate) error {
	if err := storage.DB.Create(data).Error; err != nil {
		return err
	}

	return nil
}
