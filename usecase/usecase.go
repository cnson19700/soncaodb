package usecase

import (
	"github.com/soncaodb/repository"
	"github.com/soncaodb/usecase/auth"
)

type UseCase struct {
	//User user.IUsecase
	Auth auth.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		//User: user.New(repo),
		Auth: auth.New(repo),
	}
}
