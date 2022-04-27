package Create_Role_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateRole(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateRole Suite")
}
