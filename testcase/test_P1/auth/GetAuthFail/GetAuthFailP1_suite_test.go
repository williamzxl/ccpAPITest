package GetAuthFail_test

import (
	"fmt"
	"github.com/onsi/ginkgo/reporters"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetAuthFailP1(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "GetAuthFail Suite")
	junitReport := reporters.NewJUnitReporter(fmt.Sprintf("junit_%s.xml", "GetAuthFailP1_Suite"))
	RunSpecsWithCustomReporters(t, "GetAuthFail P1 Suite", []Reporter{junitReport})
}
