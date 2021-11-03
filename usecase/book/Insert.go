package book

import (
	"context"
	"time"

	"github.com/soncaodb/model"
	"github.com/soncaodb/util/myerror"
)

type InsertRequest struct {
	Title         string    `json:"title"`
	Language      string    `json:"language"`
	Image         string    `json:"image"`
	Description   string    `json:"description"`
	ReleaseDate   time.Time `json:"release_date"`
	RatingAverage float64   `json:"rating_average"`
	TotalPage     int64     `json:"total_page"`
}

func (u *Usecase) Insert(ctx context.Context, req InsertRequest) (*model.Book, error) {
	book := &model.Book{
		Title:         req.Title,
		Image:         "https://media.istockphoto.com/photos/many-hardbound-books-background-selective-focus-picture-id1209683444?s=612x612",
		Language:      req.Language,
		Description:   req.Description,
		ReleaseDate:   req.ReleaseDate,
		RatingAverage: req.RatingAverage,
		TotalPage:     req.TotalPage,
	}

	res, err := u.bookRepo.Insert(ctx, book)

	if err != nil {
		return nil, myerror.ErrInsertBook(err)
	}
	return res, nil
}
