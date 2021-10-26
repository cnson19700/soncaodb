package myerror

import (
	"net/http"

	"github.com/cnson19700/pkg/apperror"
)

func ErrEmailFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100001,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Email Format is not valid",
		Message:   "Email Format is not valid",
	}
}

func ErrFullNameFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100001,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Full Name Format is not valid",
		Message:   "Full Name Format is not valid",
	}
}

func ErrPasswordFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100001,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Password Format is not valid",
		Message:   "Password Format is not valid",
	}
}

func ErrUserFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100001,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Create user is not valid",
		Message:   "Create user is not valid",
	}
}
