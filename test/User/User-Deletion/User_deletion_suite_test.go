package User_deletion_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserDeletion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserDeletion Suite")
}
