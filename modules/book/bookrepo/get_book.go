package bookrepo

import (
	"context"
	"versla/modules/book/bookmodel"
)

type GetBookStorage interface {
	GetBookByCondition(ctx context.Context, cond map[string]interface{}) (*bookmodel.Book, error)
}

type getBookRepo struct {
	store GetBookStorage
}

func NewGetBookRepo(store GetBookStorage) *getBookRepo {
	return &getBookRepo{
		store: store,
	}
}

func (repo *getBookRepo) GetBookDetail(ctx context.Context, id uint32) (*bookmodel.Book, error) {
	book, err := repo.store.GetBookByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return book, nil
}
