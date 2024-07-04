package config

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/internal/delivery/http/route"
)

type validationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func NewApp(route *route.ConfigRoute) *gin.Engine {
	var app = gin.New()

	route.Setup(app)

	return app
}
