package user

import (
	"github.com/soncaodb/repository"
	"github.com/soncaodb/repository/user"
)

type Usecase struct {
	userRepo user.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userRepo: repo.User,
	}
}
