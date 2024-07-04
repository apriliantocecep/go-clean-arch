package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewGin(config *viper.Viper) *gin.Engine {
	return gin.New()
}
