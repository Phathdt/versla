package bookhdl

import (
	"context"
	"versla/modules/book/bookmodel"
	"versla/sdk/sdkcm"
)

type ListBookRepo interface {
	ListBook(
		ctx context.Context,
		filter *bookmodel.ListFilter,
		paging *sdkcm.Paging,
	) ([]bookmodel.Book, error)
}

type listBookHdl struct {
	repo ListBookRepo
}

func NewListBookHdl(repo ListBookRepo) *listBookHdl {
	return &listBookHdl{repo: repo}
}

func (hdl *listBookHdl) Response(
	ctx context.Context,
	filter *bookmodel.ListFilter,
	paging *sdkcm.Paging,
) ([]bookmodel.Book, error) {
	return hdl.repo.ListBook(ctx, filter, paging)
}
