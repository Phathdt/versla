package bookstorage

import (
	"context"
	"versla/modules/book/bookmodel"
)

func (storage bookSQLStorage) DeleteBook(ctx context.Context, id uint32) error {
	if err := storage.DB.Table(bookmodel.Book{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
