package main

import (
	"fmt"
	"go-clean-arch/internal/config"
	"go-clean-arch/internal/server"
	"log"
)

func main() {
	viperConfig := config.NewViper()

	appPort := viperConfig.GetInt("APP_PORT")
	handler := server.InitializeServer()

	err := handler.Run(fmt.Sprintf(":%d", appPort))
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
