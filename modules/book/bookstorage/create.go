package bookstorage

import (
	"context"
	"versla/modules/book/bookmodel"
)

func (storage bookSQLStorage) CreateBook(ctx context.Context, data *bookmodel.BookCreate) (uint32, error) {
	if err := storage.DB.Create(data).Error; err != nil {
		return 0, err
	}

	return data.ID, nil
}
