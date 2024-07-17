package controllers

import (
	"net/http"
	"strconv"

	"myproject/backend/internal/models"
	"myproject/backend/internal/repositories"

	"github.com/labstack/echo/v4"
)

// UserController handles user-related requests
type UserController struct {
    repo repositories.UserRepository
}

// NewUserController creates a new UserController

func NewUserController(repo repositories.UserRepository) *UserController {
    return &UserController{repo: repo}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (c *UserController) GetUsers(ctx echo.Context) error {
    users, err := c.repo.GetAll()
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
    }
    return ctx.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (c *UserController) GetUser(ctx echo.Context) error {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }

    user, err := c.repo.GetByID(id)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
    }

    if user == nil {
        return ctx.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    return ctx.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (c *UserController) CreateUser(ctx echo.Context) error {
    var user models.User
    if err := ctx.Bind(&user); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }

    err := c.repo.Create(&user)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
    }

    return ctx.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(ctx echo.Context) error {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }

    var user models.User
    if err := ctx.Bind(&user); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }

    user.ID = id

    err = c.repo.Update(&user)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
    }

    return ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(ctx echo.Context) error {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }

    err = c.repo.Delete(id)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
    }

    return ctx.NoContent(http.StatusNoContent)
}
