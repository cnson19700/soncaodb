package comment

import (
	"errors"

	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/soncaodb/model"
	"github.com/soncaodb/usecase/comment"
)

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx       = &utils.CustomEchoContext{Context: c}
		appError  = apperror.AppError{}
		paginator = model.Paginator{}
		filter    = model.CommentFilter{}
	)

	if err := c.Bind(&paginator); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	if err := c.Bind(&filter); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	req := comment.GetListRequest{
		Filter:    &filter,
		Paginator: &paginator,
	}

	res, err := r.commentUseCase.GetList(ctx, &req)
	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}
	return utils.Response.Success(c, res)
}
