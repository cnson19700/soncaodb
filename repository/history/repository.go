package history

import (
	"context"

	"github.com/soncaodb/model"
)

type Repository interface {
	GetListBookByUserID(context context.Context, userID int64) ([]int64, error)
	Create(ctx context.Context, user *model.BookHistory) (*model.BookHistory, error)
}
