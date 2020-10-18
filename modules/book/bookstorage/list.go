package bookstorage

import (
	"context"
	"versla/modules/book/bookmodel"
	"versla/sdk/sdkcm"
)

func (storage bookSQLStorage) ListBooksByCondition(
	ctx context.Context,
	filter *bookmodel.ListFilter,
	paging *sdkcm.Paging,
) ([]bookmodel.Book, error) {
	var data []bookmodel.Book

	db := storage.DB.Table(bookmodel.Book{}.TableName()).
		Where("deleted_at IS NULL")

	if f := filter; f != nil {
		if v := f.Author; v != "" {
			db = db.Where("author = ?", v)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	db = db.Limit(paging.Limit).Order("id desc")

	db = db.Offset((paging.Page - 1) * paging.Limit)

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	paging.HasNext = len(data) >= paging.Limit

	return data, nil
}
