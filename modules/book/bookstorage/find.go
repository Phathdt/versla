package bookstorage

import (
	"context"
	"versla/modules/book/bookmodel"
)

func (storage *bookSQLStorage) GetBookByCondition(ctx context.Context,
	cond map[string]interface{},
) (*bookmodel.Book, error) {
	var data bookmodel.Book

	if err := storage.DB.Table(bookmodel.Book{}.TableName()).Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
