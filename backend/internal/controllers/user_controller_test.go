package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"myproject/backend/internal/controllers"
	"myproject/backend/internal/models"
	"myproject/backend/internal/repositories"
)

var _ = Describe("UserController", func() {
    var (
        mockRepo    repositories.UserRepository
        controller  *controllers.UserController
        e           *echo.Echo
        req         *http.Request
        rec         *httptest.ResponseRecorder
    )

    BeforeEach(func() {
        mockRepo = new(repositories.MockUserRepository)
        controller = controllers.NewUserController(mockRepo)
        e = echo.New()
    })

    Describe("GetUsers", func() {
        BeforeEach(func() {
            req = httptest.NewRequest(http.MethodGet, "/users", nil)
            rec = httptest.NewRecorder()
        })

        Context("when there are users", func() {
            It("should return a list of users", func() {
                mockUsers := []models.User{{ID: 1, UserName: "testuser"}}
                mockRepo.(*repositories.MockUserRepository).On("GetAll").Return(mockUsers, nil)

                c := e.NewContext(req, rec)
                err := controller.GetUsers(c)

                Expect(err).To(BeNil())
                Expect(rec.Code).To(Equal(http.StatusOK))

                var users []models.User
                json.Unmarshal(rec.Body.Bytes(), &users)
                Expect(users).To(Equal(mockUsers))
            })
        })
    })

    Describe("CreateUser", func() {
        BeforeEach(func() {
            user := &models.User{UserName: "newuser", FirstName: "New", LastName: "User", Email: "new@example.com", UserStatus: "A"}
            jsonBytes, _ := json.Marshal(user)
            req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(jsonBytes))
            req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
            rec = httptest.NewRecorder()
        })

        It("should create a new user", func() {
            mockRepo.(*repositories.MockUserRepository).On("Create", mock.AnythingOfType("*models.User")).Return(nil)

            c := e.NewContext(req, rec)
            err := controller.CreateUser(c)

            Expect(err).To(BeNil())
            Expect(rec.Code).To(Equal(http.StatusCreated))

            var createdUser models.User
            json.Unmarshal(rec.Body.Bytes(), &createdUser)
            Expect(createdUser.UserName).To(Equal("newuser"))
        })
    })
})
