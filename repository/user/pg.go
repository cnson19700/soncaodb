package user

import (
	"context"
	"log"

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

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.User, error) {
	db := r.getClient(ctx)
	user := &model.User{}

	err := db.Where("id = ?", ID).
		First(user).Error

	if err != nil {
		return nil, errors.Wrap(err, "get user by id")
	}

	return user, nil
}

func (r *pgRepository) GetAll(ctx context.Context) ([]model.User, error) {
	db := r.getClient(ctx)
	listUser := []model.User{}

	db.Find(&listUser)

	return listUser, nil
}

func (r *pgRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	db := r.getClient(ctx)
	err := db.Create(user).Error

	return user, errors.Wrap(err, "create user")
}

func (r *pgRepository) Delete(ctx context.Context, ID int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", ID).Delete(&model.User{}).Error

	if err != nil {
		log.Fatal(err)
	}

	return errors.Wrap(err, "delete fail")
}

func (r *pgRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	db := r.getClient(ctx)

	err := db.Model(&user).Updates(&user).Error

	return user, errors.Wrap(err, "update user")
}
