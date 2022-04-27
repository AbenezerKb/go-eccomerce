package Delete_Store_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeleteStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeleteStore Suite")
}
