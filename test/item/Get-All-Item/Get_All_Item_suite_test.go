package Get_All_Item_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetAllItem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetAllItem Suite")
}
