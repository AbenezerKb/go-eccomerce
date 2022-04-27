package Update_Order_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUpdateOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UpdateOrder Suite")
}
