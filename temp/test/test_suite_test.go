package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

func RegisterFailHandler(Fail2 func(message string, callerSkip ...int)) {

}
