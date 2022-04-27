package User_Login_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserLogin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserLogin Suite")
}
