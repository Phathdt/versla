package bookrepo

import (
	"context"
	"versla/modules/book/bookmodel"
)

type CreateBookStorage interface {
	CreateBook(ctx context.Context, data *bookmodel.BookCreate) (uint32, error)
	GetBookByCondition(ctx context.Context, cond map[string]interface{}) (*bookmodel.Book, error)
}

type createBookRepo struct {
	store CreateBookStorage
}

func NewCreateBookRepo(store CreateBookStorage) *createBookRepo {
	return &createBookRepo{
		store: store,
	}
}

func (repo *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.BookCreate) (*bookmodel.Book, error) {
	bookID, err := repo.store.CreateBook(ctx, data)

	if err != nil {
		return nil, err
	}

	book, err := repo.store.GetBookByCondition(ctx, map[string]interface{}{"id": bookID})

	if err != nil {
		return nil, err
	}

	return book, nil
}
