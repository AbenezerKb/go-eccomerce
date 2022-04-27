package Update_Store_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUpdateStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UpdateStore Suite")
}
