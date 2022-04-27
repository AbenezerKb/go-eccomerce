package Get_All_Store_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetAllStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetAllStore Suite")
}
