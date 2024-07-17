package repositories_test

import (
	"database/sql"
	"myproject/backend/internal/models"
	"myproject/backend/internal/repositories"
	"os"

	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserRepository", func() {
    var (
        db   *sql.DB
        repo repositories.UserRepository
    )

    BeforeEach(func() {
        // Set up the database connection
        connStr := os.Getenv("DATABASE_URL")
		if connStr == "" {
			connStr = "postgres://postgres:9441@localhost:5432/postgres?sslmode=disable"
		}
		var err error
		db, err = sql.Open("postgres", connStr)
		Expect(err).NotTo(HaveOccurred())

        // Create a new UserRepository instance
        repo = repositories.NewUserRepository(db)

        // Clear the users table before each test
        _, err = db.Exec("DELETE FROM users")
        Expect(err).NotTo(HaveOccurred())
    })

    AfterEach(func() {
        db.Close()
    })

    Describe("Create", func() {
        It("should create a new user", func() {
            user := &models.User{
                UserName:   "testuser",
                FirstName:  "Test",
                LastName:   "User",
                Email:      "test@example.com",
                UserStatus: "A",
                Department: "IT",
            }

            err := repo.Create(user)
            Expect(err).NotTo(HaveOccurred())
            Expect(user.ID).To(BeNumerically(">", 0))
        })
    })

    Describe("GetByID", func() {
        It("should retrieve a user by ID", func() {
            user := &models.User{
                UserName:   "testuser",
                FirstName:  "Test",
                LastName:   "User",
                Email:      "test@example.com",
                UserStatus: "A",
                Department: "IT",
            }

            err := repo.Create(user)
            Expect(err).NotTo(HaveOccurred())

            retrievedUser, err := repo.GetByID(user.ID)
            Expect(err).NotTo(HaveOccurred())
            Expect(retrievedUser).To(Equal(user))
        })

        It("should return nil if user not found", func() {
            user, err := repo.GetByID(999) // Assuming 999 is not a valid user ID
            Expect(err).NotTo(HaveOccurred())
            Expect(user).To(BeNil())
        })
    })

    Describe("Update", func() {
        It("should update an existing user", func() {
            user := &models.User{
                UserName:   "testuser",
                FirstName:  "Test",
                LastName:   "User",
                Email:      "test@example.com",
                UserStatus: "A",
                Department: "IT",
            }

            err := repo.Create(user)
            Expect(err).NotTo(HaveOccurred())

            user.FirstName = "Updated"
            user.LastName = "Name"

            err = repo.Update(user)
            Expect(err).NotTo(HaveOccurred())

            updatedUser, err := repo.GetByID(user.ID)
            Expect(err).NotTo(HaveOccurred())
            Expect(updatedUser.FirstName).To(Equal("Updated"))
            Expect(updatedUser.LastName).To(Equal("Name"))
        })
    })

    Describe("Delete", func() {
        It("should delete an existing user", func() {
            user := &models.User{
                UserName:   "testuser",
                FirstName:  "Test",
                LastName:   "User",
                Email:      "test@example.com",
                UserStatus: "A",
                Department: "IT",
            }

            err := repo.Create(user)
            Expect(err).NotTo(HaveOccurred())

            err = repo.Delete(user.ID)
            Expect(err).NotTo(HaveOccurred())

            deletedUser, err := repo.GetByID(user.ID)
            Expect(err).NotTo(HaveOccurred())
            Expect(deletedUser).To(BeNil())
        })
    })
})
