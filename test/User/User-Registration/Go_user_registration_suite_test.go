package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoUserRegistration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoUserRegistration Suite")
}
