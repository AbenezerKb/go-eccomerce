package User_Update_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserUpdate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserUpdate Suite")
}
