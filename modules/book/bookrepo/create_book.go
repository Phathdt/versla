package bookrepo

import (
	"context"
	"versla/modules/book/bookmodel"
)

type CreateBookStorage interface {
	CreateBook(ctx context.Context, data *bookmodel.BookCreate) (uint, error)
}

type createBookRepo struct {
	store CreateBookStorage
}

func NewCreateBookRepo(store CreateBookStorage) *createBookRepo {
	return &createBookRepo{
		store: store,
	}
}

func (repo *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.BookCreate) (uint, error) {
	bookID, err := repo.store.CreateBook(ctx, data)

	if err != nil {
		return 0, err
	}

	return bookID, nil
}
