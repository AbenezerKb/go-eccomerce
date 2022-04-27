package Get_All_Orders_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetAllOrders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetAllOrders Suite")
}
