package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-clean-arch/internal/helper"
	"go-clean-arch/internal/model"
	"go-clean-arch/internal/usecase"
	"net/http"
)

type AuthController struct {
	UseCase *usecase.UserUseCase
	Log     *logrus.Logger
}

func NewAuthController(useCase *usecase.UserUseCase, log *logrus.Logger) *AuthController {
	return &AuthController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var registerRequest model.RegisterUserRequest

	if err := ctx.ShouldBindBodyWithJSON(&registerRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.UseCase.Register(ctx, &registerRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse(ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, res)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var loginRequest model.LoginUserRequest

	if err := ctx.ShouldBindBodyWithJSON(&loginRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.UseCase.Login(ctx, &loginRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse(ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, res)
}
