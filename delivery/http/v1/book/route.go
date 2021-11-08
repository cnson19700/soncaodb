package book

import (
	"github.com/labstack/echo/v4"
	"github.com/soncaodb/usecase"
	"github.com/soncaodb/usecase/book"
)

type Route struct {
	bookUseCase book.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		bookUseCase: useCase.Book,
	}
	group.POST("", r.Insert)
	group.DELETE("/:id", r.Delete)
	group.PUT("/:id", r.Update)
	group.GET("", r.SearchBook)
	group.GET("/find", r.GetBook)
}
