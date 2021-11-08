package comment

import (
	"context"

	"github.com/soncaodb/model"
)

type Repository interface {
	GetById(ctx context.Context, ID int64) (*model.Comment, error)
	GetAll(ctx context.Context) ([]model.Comment, error)
	Delete(ctx context.Context, ID int64) error
	DeleteSubComment(ctx context.Context, parentID int64) error
	Insert(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Update(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	// SearchComment(ctx context.Context,
	// 	page *model.Paginator,
	// 	filter *model.CommentFilter,
	// 	orders []string) (*model.CommentResult, error)
}
