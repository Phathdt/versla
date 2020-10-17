package bookstorage

import (
	"context"
	"versla/modules/book/bookmodel"
)

func (storage *bookSQLStorage) GetBookByCondition(ctx context.Context,
	cond map[string]interface{},
) (*bookmodel.Book, error) {
	var data bookmodel.Book

	err := storage.DB.Where(cond).First(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}
