package history

import (
	"context"

	"github.com/pkg/errors"
	"github.com/soncaodb/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getClient}
}

func (r *pgRepository) GetListBookByUserID(ctx context.Context, userID int64) ([]int64, error) {
	db := r.getClient(ctx)
	var history []model.BookHistory
	var res = []int64{}

	err := db.Where("user_id = ? ", userID).
		Find(&history).Error

	if err != nil {
		return res, errors.Wrap(err, "book history")
	}

	for i := 0; i < len(history); i++ {
		res = append(res, history[i].BookID)
	}

	return res, nil
}

func (r *pgRepository) Create(ctx context.Context, user *model.BookHistory) (*model.BookHistory, error) {
	db := r.getClient(ctx)
	err := db.Create(user).Error

	return user, errors.Wrap(err, "create user")
}
