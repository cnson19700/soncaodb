package comment

import (
	"github.com/soncaodb/repository"
	"github.com/soncaodb/repository/comment"
)

type Usecase struct {
	commentRepo comment.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		commentRepo: repo.Comment,
	}
}


