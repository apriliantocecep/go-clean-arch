package route

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/internal/delivery/http/controller"
	"go-clean-arch/internal/delivery/http/middleware"
)

type ConfigRoute struct {
	AuthMiddleware *middleware.AuthMiddleware
	AuthController *controller.AuthController
	UserController *controller.UserController
}

func NewRoute(
	authMiddleware *middleware.AuthMiddleware,
	authController *controller.AuthController,
	userController *controller.UserController,
) *ConfigRoute {
	return &ConfigRoute{
		AuthMiddleware: authMiddleware,
		AuthController: authController,
		UserController: userController,
	}
}

func (c *ConfigRoute) Setup(app *gin.Engine) {
	c.guestApiRoute(app)
	c.protectedApiRoute(app)
}

func (c *ConfigRoute) guestApiRoute(app *gin.Engine) {
	api := apiGroup(app)
	{
		api.POST("/register", c.AuthController.Register)
		api.POST("/login", c.AuthController.Login)
	}
}

func (c *ConfigRoute) protectedApiRoute(app *gin.Engine) {
	api := apiGroup(app).Use(c.AuthMiddleware.TokenAuthorization)
	{
		api.GET("/me", c.UserController.GetUser)
	}
}

func apiGroup(app *gin.Engine) *gin.RouterGroup {
	return app.Group("/api")
}
