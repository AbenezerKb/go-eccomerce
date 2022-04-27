package Get_Store_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetStore Suite")
}
