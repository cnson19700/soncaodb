package auth

import (
	"context"
	"log"

	"github.com/soncaodb/model"
	"github.com/soncaodb/package/auth"
	checkform "github.com/soncaodb/package/checkForm"
	"github.com/soncaodb/util/myerror"
)

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func (u *Usecase) Register(ctx context.Context, req RegisterRequest) (*model.User, error) {
	//Email Format Error
	isMail, email := checkform.CheckFormatValue("email", req.Email)
	if !isMail {
		return &model.User{}, myerror.ErrEmailFormat(nil)
	}

	//password format error
	isPass, password := checkform.CheckFormatValue("password", req.Password)
	if !isPass {
		return &model.User{}, myerror.ErrEmailFormat(nil)
	}

	//password format error
	isFullName, fullname := checkform.CheckFormatValue("full_name", req.FullName)
	if !isFullName {
		return &model.User{}, myerror.ErrEmailFormat(nil)
	}

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	var user = &model.User{
		FullName: fullname,
		Password: passwordHash,
		Age:      req.Age,
		Email:    email,
	}

	res, err := u.userRepo.Create(ctx, user)

	if err != nil {
		return &model.User{}, nil
	}
	return res, nil
}
