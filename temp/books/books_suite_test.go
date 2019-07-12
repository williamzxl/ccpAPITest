package books_test

import (
	. "github.com/onsi/ginkgo"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

func RegisterFailHandler(Fail2 func(message string, callerSkip ...int)) {

}
