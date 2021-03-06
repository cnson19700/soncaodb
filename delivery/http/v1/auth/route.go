package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/soncaodb/usecase"
	"github.com/soncaodb/usecase/auth"
)

type Route struct {
	authUseCase auth.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		authUseCase: useCase.Auth,
	}

	group.POST("/register", r.Register)
	group.POST("/login", r.Login)
}
