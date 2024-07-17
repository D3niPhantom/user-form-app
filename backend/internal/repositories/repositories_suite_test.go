package repositories_test

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepositories(t *testing.T) {
    RegisterFailHandler(Fail)
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Println("Warning: .env file not found. Using environment variables.")
    }
    RunSpecs(t, "Repositories Suite")
}
