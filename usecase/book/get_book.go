package book

import (
	"context"

	"github.com/soncaodb/model"
	"github.com/soncaodb/util/myerror"
)

type GetBookRequest struct {
	ID int64 `json:"id"`
}

func (u *Usecase) GetBook(ctx context.Context, req GetBookRequest) (*model.Book, error) {
	book, err := u.bookRepo.GetById(ctx, req.ID)
	if err != nil {
		return nil, myerror.ErrGetBook(err)
	}

	return book, err
}
