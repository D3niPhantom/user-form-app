package routes_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"myproject/backend/internal/controllers"
	"myproject/backend/internal/models"
	"myproject/backend/internal/repositories"
	"myproject/backend/internal/routes"
)

var _ = Describe("Routes", func() {
    var (
        e           *echo.Echo
        mockRepo    *repositories.MockUserRepository
        userController *controllers.UserController
    )

    BeforeEach(func() {
        e = echo.New()
        mockRepo = new(repositories.MockUserRepository)
        userController = controllers.NewUserController(mockRepo)
        routes.SetupRoutes(e, userController)
    })

    Describe("GET /api/users", func() {
        It("should return a list of users", func() {
            mockUsers := []models.User{{ID: 1, UserName: "testuser"}}
            mockRepo.On("GetAll").Return(mockUsers, nil)

            req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
            rec := httptest.NewRecorder()
            e.ServeHTTP(rec, req)

            Expect(rec.Code).To(Equal(http.StatusOK))

            var users []models.User
            json.Unmarshal(rec.Body.Bytes(), &users)
            Expect(users).To(Equal(mockUsers))
        })
    })

    Describe("POST /api/users", func() {
        It("should create a new user", func() {
            user := &models.User{UserName: "newuser", FirstName: "New", LastName: "User", Email: "new@example.com", UserStatus: "A"}
            mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

            jsonBody, _ := json.Marshal(user)
            req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(jsonBody)))
            req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
            rec := httptest.NewRecorder()
            e.ServeHTTP(rec, req)

            Expect(rec.Code).To(Equal(http.StatusCreated))

            var createdUser models.User
            json.Unmarshal(rec.Body.Bytes(), &createdUser)
            Expect(createdUser.UserName).To(Equal("newuser"))
        })
    })
})
