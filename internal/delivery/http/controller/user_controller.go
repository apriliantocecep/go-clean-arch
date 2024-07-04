package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-clean-arch/internal/helper"
	"go-clean-arch/internal/model"
	"go-clean-arch/internal/usecase"
	"net/http"
)

type UserController struct {
	*usecase.UserUseCase
	Log *logrus.Logger
}

func NewUserController(userUseCase *usecase.UserUseCase, log *logrus.Logger) *UserController {
	return &UserController{UserUseCase: userUseCase, Log: log}
}

func (u *UserController) GetUser(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	res, err := u.UserUseCase.GetUser(ctx, &model.GetUserRequest{ID: userId.(uint)})
	if err != nil {
		helper.ErrorResponse(ctx, err.Code, err.Details)
		ctx.Abort()
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, res)
}
