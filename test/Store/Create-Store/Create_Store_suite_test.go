package Create_Store_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateStore Suite")
}
