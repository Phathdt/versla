package bookhdl

import (
	"context"
	"versla/modules/book/bookmodel"
)

type UpdateBookRepo interface {
	UpdateBook(ctx context.Context, id uint32, data *bookmodel.BookUpdate) (*bookmodel.Book, error)
}

type updateBookHdl struct {
	repo UpdateBookRepo
}

func NewUpdateBookHdl(repo UpdateBookRepo) *updateBookHdl {
	return &updateBookHdl{
		repo: repo,
	}
}

func (hdl *updateBookHdl) Response(ctx context.Context, id uint32, data *bookmodel.BookUpdate) (*bookmodel.Book, error) {
	return hdl.repo.UpdateBook(ctx, id, data)
}
