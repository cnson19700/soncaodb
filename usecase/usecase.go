package usecase

import (
	"github.com/soncaodb/repository"
	"github.com/soncaodb/usecase/auth"
	"github.com/soncaodb/usecase/book"
	"github.com/soncaodb/usecase/user"
)

type UseCase struct {
	User user.IUsecase
	Auth auth.IUsecase
	Book book.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User: user.New(repo),
		Auth: auth.New(repo),
		Book: book.New(repo),
	}
}
