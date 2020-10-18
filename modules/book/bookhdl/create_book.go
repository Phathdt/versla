package bookhdl

import (
	"context"
	"versla/modules/book/bookmodel"
)

type CreateBookRepo interface {
	CreateBook(ctx context.Context, data *bookmodel.BookCreate) (uint, error)
}

type createBookHdl struct {
	repo CreateBookRepo
}

func NewCreateBookHdl(repo CreateBookRepo) *createBookHdl {
	return &createBookHdl{
		repo: repo,
	}
}

func (hdl *createBookHdl) Response(ctx context.Context, data *bookmodel.BookCreate) (uint, error) {
	return hdl.repo.CreateBook(ctx, data)
}
