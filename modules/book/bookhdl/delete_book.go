package bookhdl

import "context"

type DeleteBookRepo interface {
	DeleteBook(ctx context.Context, id uint32) error
}

type deleteBookRepo struct {
	store DeleteBookRepo
}

func NewDeleteBookHdl(store DeleteBookRepo) *deleteBookRepo {
	return &deleteBookRepo{
		store: store,
	}
}

func (hdl *deleteBookRepo) Response(ctx context.Context, id uint32) error {
	return hdl.store.DeleteBook(ctx, id)
}
