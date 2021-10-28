package myerror

import (
	"net/http"

	"github.com/cnson19700/pkg/apperror"
)

func ErrUserFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100001,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Create user is not valid",
		Message:   "Create user is not valid",
	}
}

func ErrGetUser(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to get User",
		Message:   "Fail to get USer",
	}
}

func ErrUpdatePassword(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to update password",
		Message:   "Fail to update password",
	}
}

func ErrInvalidPassword(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100020,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Password is not correct",
		Message:   "Password is not correct",
	}
}

func ErrMatchPassword(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100030,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "2 Password is not match",
		Message:   "2 Password is not match",
	}
}
