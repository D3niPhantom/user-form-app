package main

import (
	"database/sql"
	"log"
	"myproject/backend/internal/controllers"
	"myproject/backend/internal/repositories"
	"myproject/backend/internal/routes"

	_ "myproject/backend/docs" // Generated Swag Docs

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title User Management API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:8080
// @BasePath /api
func main() {
    // modify this connection string to match your postgres settings.
    dbConnectionString := "postgres://username:password@localhost:5432/database_name?sslmode=disable"
    db, err := sql.Open("postgres", dbConnectionString)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    userRepo := repositories.NewUserRepository(db)
    userController := controllers.NewUserController(userRepo)

    routes.SetupRoutes(e, userController)

    // Add Swagger route
    e.GET("/swagger/*", echoSwagger.WrapHandler)

    e.Logger.Fatal(e.Start(":8080"))
}
