package config

import (
	"github.com/spf13/viper"
	"go-clean-arch/internal/model"
)

func NewJwtWrapper(config *viper.Viper) *model.JwtWrapper {
	return &model.JwtWrapper{
		SecretKey:       config.GetString("JWT_SECRET_KEY"),
		Issuer:          config.GetString("JWT_ISSUER"),
		ExpirationHours: config.GetInt64("JWT_EXPIRATION_HOURS"),
	}
}
