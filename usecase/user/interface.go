package user

import (
	"context"

	"github.com/soncaodb/model"
)

type IUsecase interface {
	GetMe(ctx context.Context) (*model.User, error)
	UpdatePassword(ctx context.Context, req UpdatePasswordRequest) error
}
