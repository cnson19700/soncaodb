package student

import (
	"context"

	"github.com/soncaodb/model"
)

type Repository interface {
	GetById(ctx context.Context, ID int64) (*model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
	Delete(ctx context.Context, ID int64) error
	Create(ctx context.Context, student *model.Student) (*model.Student, error)
	ShowStudent(ctx context.Context, st[]model.Student)
	Update(ctx context.Context, student *model.Student) (*model.Student, error)
}
