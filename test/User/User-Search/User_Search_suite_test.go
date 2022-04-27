package User_Search_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserSearch Suite")
}
