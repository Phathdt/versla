package bookrepo

import (
	"context"
	"versla/modules/book/bookmodel"
	"versla/sdk/sdkcm"
)

type ListBookStorage interface {
	ListBooksByCondition(
		ctx context.Context,
		filter *bookmodel.ListFilter,
		paging *sdkcm.Paging,
	) ([]bookmodel.Book, error)
}

type listBookRepo struct {
	store ListBookStorage
}

func NewListBookRepo(store ListBookStorage) *listBookRepo {
	return &listBookRepo{
		store: store,
	}
}

func (repo *listBookRepo) ListBook(
	ctx context.Context,
	filter *bookmodel.ListFilter,
	paging *sdkcm.Paging,
) ([]bookmodel.Book, error) {
	result, err := repo.store.ListBooksByCondition(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
