package Create_Item_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateItem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateItem Suite")
}
