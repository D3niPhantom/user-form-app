package repositories

import "myproject/backend/internal/models"

type UserRepository interface {
    GetAll() ([]models.User, error)
    GetByID(id int) (*models.User, error)
    Create(user *models.User) error
    Update(user *models.User) error
    Delete(id int) error
}
