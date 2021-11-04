package book

import (
	"context"

	"github.com/soncaodb/model"
)

type IUsecase interface {
	Insert(ctx context.Context, req InsertRequest) (*model.Book, error)
	Delete(ctx context.Context, req DeleteBookRequest) error
	Update(ctx context.Context, req UpdateBookRequest) (*model.Book, error)
	SearchBook(ctx context.Context, searchText string, req SearchBookRequest) (*model.BookResult, error)
}
