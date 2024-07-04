//go:build wireinject
// +build wireinject

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-clean-arch/internal/config"
	"go-clean-arch/internal/delivery/http/controller"
	"go-clean-arch/internal/delivery/http/middleware"
	"go-clean-arch/internal/delivery/http/route"
	"go-clean-arch/internal/repository"
	"go-clean-arch/internal/usecase"
)

var configSet = wire.NewSet(config.NewViper, config.NewLogger, config.NewDatabase, config.NewValidator, config.NewJwtWrapper)
var repositorySet = wire.NewSet(repository.NewUserRepository)
var useCaseSet = wire.NewSet(usecase.NewUserUseCase)
var controllerSet = wire.NewSet(controller.NewAuthController, controller.NewUserController)
var middlewareSet = wire.NewSet(middleware.NewAuthMiddleware)

func InitializeServer() *gin.Engine {
	wire.Build(configSet, repositorySet, useCaseSet, controllerSet, middlewareSet, route.NewRoute, config.NewApp)
	return nil
}
