package user

import (
	"context"

	"github.com/cnson19700/pkg/middleware"
	"github.com/soncaodb/model"
	"github.com/soncaodb/util/myerror"
)

func (u *Usecase) GetMe(ctx context.Context) (*model.User, error) {
	payload := middleware.GetClaim(ctx)
	id := payload.UserID

	user, err := u.userRepo.GetById(ctx, id)
	if err != nil {
		return &model.User{}, myerror.ErrGetUser(err)
	}

	return user, nil
}
