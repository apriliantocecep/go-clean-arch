package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-clean-arch/internal/model"
	apperrors "go-clean-arch/pkg/errors"
	"net/http"
	"strings"
)

func SuccessResponse(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, model.ApiResponse{Data: data})
}

func ErrorResponse(ctx *gin.Context, code int, details any) {
	ctx.AbortWithStatusJSON(code, model.ApiResponse{Details: details})
}

func ErrorBindingResponse(ctx *gin.Context, validationErrors validator.ValidationErrors) {
	var responseErrors []model.ErrorDetail
	for _, fieldError := range validationErrors {
		field := fieldError.Field()
		responseErrors = append(responseErrors, model.ErrorDetail{
			Field:          strings.ToLower(field),
			ValidationCode: strings.ToUpper(fieldError.Tag()),
			Param:          fieldError.Param(),
			Message:        ParseFieldError(fieldError),
		})
	}
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ApiResponse{Details: responseErrors})
}

func ValidationErrorResponse(ctx *gin.Context, err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		ErrorBindingResponse(ctx, validationErrors)
	} else {
		var outErr []model.ErrorDetail
		outErr = append(outErr, model.ErrorDetail{
			Field:          "general",
			ValidationCode: string(apperrors.InvalidJson),
			Param:          "",
			Message:        err.Error(),
		})
		ErrorResponse(ctx, http.StatusBadRequest, outErr)
	}
}
