package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.AddConfigPath("./configs/")
	config.AddConfigPath("./")
	config.SetConfigName(".env")
	config.SetConfigType("env")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}
