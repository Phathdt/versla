package bookrepo

import (
	"context"
	"versla/modules/book/bookmodel"
)

type UpdateBookStorage interface {
	GetBookByCondition(ctx context.Context, cond map[string]interface{}) (*bookmodel.Book, error)
	UpdateBook(ctx context.Context, id uint32, data *bookmodel.BookUpdate) error
}

type updateBookRepo struct {
	store UpdateBookStorage
}

func NewUpdateBookRepo(store UpdateBookStorage) *updateBookRepo {
	return &updateBookRepo{
		store: store,
	}
}

func (repo *updateBookRepo) UpdateBook(ctx context.Context, id uint32, data *bookmodel.BookUpdate) (*bookmodel.Book, error) {
	_, err := repo.store.GetBookByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	if err := repo.store.UpdateBook(ctx, id, data); err != nil {
		return nil, err
	}

	book, err := repo.store.GetBookByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return book, nil
}
