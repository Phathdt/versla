package bookstorage

import (
	"context"
	"fmt"
	"versla/modules/book/bookmodel"
)

func (storage bookSQLStorage) UpdateBook(ctx context.Context, id uint32, data *bookmodel.BookUpdate) error {
	tx := storage.DB.Begin()

	if err := tx.Table(bookmodel.Book{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println("ROLLBACK")
		tx.Rollback()
		return err
	}

	return nil
}
