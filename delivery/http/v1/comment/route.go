package comment

import (
	"github.com/labstack/echo/v4"

	"github.com/soncaodb/usecase"
	"github.com/soncaodb/usecase/comment"
)

type Route struct {
	commentUseCase comment.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		commentUseCase: useCase.Comment,
	}
	group.POST("", r.Insert)
	group.DELETE("", r.Delete)
	group.GET("", r.GetList)
}
