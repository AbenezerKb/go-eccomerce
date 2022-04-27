package Get_Item_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetItem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetItem Suite")
}
