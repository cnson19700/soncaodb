package book

import (
	"context"

	"github.com/soncaodb/model"
)

type Repository interface {
	GetById(ctx context.Context, ID int64) (*model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	Delete(ctx context.Context, ID int64) error
	Insert(ctx context.Context, user *model.Book) (*model.Book, error)
	Update(ctx context.Context, user *model.Book) (*model.Book, error)
	SearchBook(ctx context.Context,
		page *model.Paginator,
		searchText string,
		filter *model.BookFilter,
		orders []string) (*model.BookResult, error)
}
