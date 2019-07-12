package GetAuthFail_test

import (
	"fmt"
	"github.com/onsi/ginkgo/reporters"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetAuthFail(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "GetAuthFail Suite")
	junitReport := reporters.NewJUnitReporter(fmt.Sprintf("junit_%s.xml", "GetAuthFail_Suite"))
	RunSpecsWithCustomReporters(t, "GetAuthFail Suite", []Reporter{junitReport})
}

//func TestGetAuthSuccess(t *testing.T) {
//	RegisterFailHandler(Fail)
//	//RunSpecs(t, "GetAuthFail Suite")
//	junitReport := reporters.NewJUnitReporter(fmt.Sprintf("junit_%s.xml", "GetAuthSuccess_Suite"))
//	RunSpecsWithCustomReporters(t, "GetAuthSuccess Suite", []Reporter{junitReport})
//}
