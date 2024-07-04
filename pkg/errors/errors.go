package errors

import "go-clean-arch/internal/model"

type AppError struct {
	Code    int
	Err     error
	Details model.ErrorDetail
}

func NewAppError(code int, details model.ErrorDetail) *AppError {
	appError := &AppError{}
	appError.Code = code
	appError.Details = details
	return appError
}

func (appErr AppError) WithError(err error) *AppError {
	appErr.Err = err
	return &appErr
}
