package comment

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

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.Comment, error) {
	db := r.getClient(ctx)
	comment := &model.Comment{}

	err := db.Where("id = ?", ID).
		First(comment).Error

	if err != nil {
		return nil, errors.Wrap(err, "get comment by id")
	}

	return comment, nil
}

func (r *pgRepository) GetAll(ctx context.Context) ([]model.Comment, error) {
	db := r.getClient(ctx)
	listComment := []model.Comment{}

	db.Find(&listComment)

	return listComment, nil
}

func (r *pgRepository) Insert(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	db := r.getClient(ctx)
	err := db.Create(comment).Error

	return comment, errors.Wrap(err, "create comment")
}

func (r *pgRepository) Delete(ctx context.Context, id int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", id).Delete(&model.Comment{}).Error

	return errors.Wrap(err, "delete comment fail")
}

func (r *pgRepository) DeleteSubComment(ctx context.Context, parentID int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", parentID).Delete(&model.Comment{}).Error

	return errors.Wrap(err, "delete comment fail")
}

func (r *pgRepository) Update(ctx context.Context, book *model.Comment) (*model.Comment, error) {
	db := r.getClient(ctx)
	err := db.Save(book).Error

	return book, errors.Wrap(err, "update comment fail")
}
