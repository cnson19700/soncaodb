package user

import (
	"github.com/labstack/echo/v4"
	"github.com/soncaodb/usecase"
	"github.com/soncaodb/usecase/user"
)

type Route struct {
	userUseCase user.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		userUseCase: useCase.User,
	}
	group.GET("/getme", r.GetMe)
	group.PUT("/updatepassword", r.UpdatePassword)
}
