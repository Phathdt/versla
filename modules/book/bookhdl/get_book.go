package bookhdl

import (
	"context"
	"versla/modules/book/bookmodel"
)

type GetBookRepo interface {
	GetBookDetail(ctx context.Context, id uint32) (*bookmodel.Book, error)
}

type getBookHdl struct {
	repo GetBookRepo
}

func NewGetBookHdl(repo GetBookRepo) *getBookHdl {
	return &getBookHdl{
		repo: repo,
	}
}

func (hdl *getBookHdl) Response(ctx context.Context, id uint32) (*bookmodel.Book, error) {
	return hdl.repo.GetBookDetail(ctx, id)
}
