package middleware

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/internal/helper"
	"go-clean-arch/internal/model"
	"go-clean-arch/internal/usecase"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	UseCase *usecase.UserUseCase
}

func NewAuthMiddleware(userUseCase *usecase.UserUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		UseCase: userUseCase,
	}
}

func (a *AuthMiddleware) TokenAuthorization(ctx *gin.Context) {
	var errOut []model.ErrorDetail
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	res, err := a.UseCase.Verify(ctx, &model.ValidateUserRequest{Token: token[1]})
	if err != nil {
		errOut = append(errOut, err.Details)
		helper.ErrorResponse(ctx, err.Code, errOut)
		ctx.Abort()
		return
	}

	ctx.Set("userId", res.User.ID)
	ctx.Next()
}
