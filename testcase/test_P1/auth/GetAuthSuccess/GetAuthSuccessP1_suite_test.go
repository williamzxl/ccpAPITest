package GetAuthSuccess_test

import (
	"fmt"
	"github.com/onsi/ginkgo/reporters"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetAuthSuccess(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "GetAuthFail Suite")
	junitReport := reporters.NewJUnitReporter(fmt.Sprintf("junit_%s.xml", "GetAuthSuccessP1_Suite"))
	RunSpecsWithCustomReporters(t, "GetAuthSuccessP1 Suite", []Reporter{junitReport})
}
