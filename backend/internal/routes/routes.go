package routes

import (
	"myproject/backend/internal/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, uc *controllers.UserController) {
	api := e.Group("/api")

	api.GET("/users", uc.GetUsers)
	api.GET("/users/:id", uc.GetUser)
	api.POST("/users", uc.CreateUser)
	api.PUT("/users/:id", uc.UpdateUser)
	api.DELETE("/users/:id", uc.DeleteUser)
}
