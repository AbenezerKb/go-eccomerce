package Get_Orders_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetOrders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetOrders Suite")
}
