package bookrepo

import "context"

type DeleteBookStorage interface {
	DeleteBook(ctx context.Context, id uint32) error
}

type deleteBookRepo struct {
	store DeleteBookStorage
}

func NewDeleteBookRepo(store DeleteBookStorage) *deleteBookRepo {
	return &deleteBookRepo{
		store: store,
	}
}

func (repo *deleteBookRepo) DeleteBook(ctx context.Context, id uint32) error {
	err := repo.store.DeleteBook(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
