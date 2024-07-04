package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ParseFieldError(e validator.FieldError) string {
	fieldPrefix := fmt.Sprintf("The %s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", fieldPrefix)
	case "email":
		return fmt.Sprintf("%s must be an email", fieldPrefix)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", fieldPrefix, e.Param())
	case "let":
		return fmt.Sprintf("%s must be less than or equal to %s", fieldPrefix, e.Param())
	default:
		return fmt.Errorf("%v", e).Error()
	}
}
