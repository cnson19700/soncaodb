package auth

import (
	"context"

	"github.com/cnson19700/pkg/middleware"
	"github.com/soncaodb/config"
	"github.com/soncaodb/model"
	"github.com/soncaodb/package/auth"
	checkform "github.com/soncaodb/package/checkForm"
	"github.com/soncaodb/util/myerror"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *Usecase) Login(ctx context.Context, req LoginRequest) (string, error) {
	//Email Format Error
	isMail, email := checkform.CheckFormatValue("email", req.Email)
	if !isMail {
		return "", myerror.ErrEmailFormat(nil)
	}

	//password format error
	isPass, password := checkform.CheckFormatValue("password", req.Password)
	if !isPass {
		return "", myerror.ErrEmailFormat(nil)
	}

	user, err := u.userRepo.GetEmail(ctx, email)
	if err != nil {
		return "", myerror.ErrLogin(err)
	}

	if (&model.User{}) == user {
		return "", myerror.ErrValidate(nil)
	}

	//check email correct
	isPassCorrect := auth.VerifyPassword(password, user.Password)
	if !isPassCorrect {
		return "", myerror.ErrHashPassword(err)
	}

	tokenService := middleware.NewTokenSvc(config.GetConfig().Jwt.Key)
	expireTime := config.GetConfig().Jwt.TokenExpire
	t, err := tokenService.Encode(user.ID, user.Email, "soncao", expireTime)
	if err != nil {
		return "", myerror.ErrToken(err)
	}
	return t, err
}
