package book

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

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.Book, error) {
	db := r.getClient(ctx)
	book := &model.Book{}

	err := db.Where("id = ?", ID).
		First(book).Error

	if err != nil {
		return nil, errors.Wrap(err, "get book by id")
	}

	return book, nil
}

func (r *pgRepository) GetAll(ctx context.Context) ([]model.Book, error) {
	db := r.getClient(ctx)
	listBook := []model.Book{}

	db.Find(&listBook)

	return listBook, nil
}

func (r *pgRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	db := r.getClient(ctx)
	err := db.Create(book).Error

	return book, errors.Wrap(err, "create book")
}

func (r *pgRepository) Delete(ctx context.Context, id int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", id).Delete(&model.Book{}).Error

	return errors.Wrap(err, "delete book fail")
}

func (r *pgRepository) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	db := r.getClient(ctx)
	err := db.Save(book).Error

	return book, errors.Wrap(err, "update book fail")
}

func (r *pgRepository) GetTitle(ctx context.Context, tile string) (*model.Book, error) {
	db := r.getClient(ctx)
	book := &model.Book{}

	err := db.Where("title = ?", tile).
		First(book).Error

	if err != nil {
		return nil, errors.Wrap(err, "get book by title")
	}

	return book, nil
}
