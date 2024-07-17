package repositories

import (
	"myproject/backend/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetAll() ([]models.User, error) {
    args := m.Called()
    return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id int) (*models.User, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *models.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserRepository) Update(user *models.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserRepository) Delete(id int) error {
    args := m.Called(id)
    return args.Error(0)
}
