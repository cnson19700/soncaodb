package book

import (
	"github.com/soncaodb/repository"
	"github.com/soncaodb/repository/book"
)

type Usecase struct {
	bookRepo book.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		bookRepo: repo.Book,
	}
}
