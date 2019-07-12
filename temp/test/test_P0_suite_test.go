package main_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	//"github.com/onsi/ginkgo/config"
	//"github.com/onsi/ginkgo/types"
	"testing"
)

//type Reporter1 interface {
//	SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary)
//	BeforeSuiteDidRun(setupSummary *types.SetupSummary)
//	SpecWillRun(specSummary *types.SpecSummary)
//	SpecDidComplete(specSummary *types.SpecSummary)
//	AfterSuiteDidRun(setupSummary *types.SetupSummary)
//	SpecSuiteDidEnd(summary *types.SuiteSummary)
//}

func TestTestP0(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "TestP0 Suite")
	junitReporter := reporters.NewJUnitReporter("junit12.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Foo Suite", []Reporter{junitReporter})
	//RunSpecsWithDefaultAndCustomReporters(t *testing.T, description string, reporters []Reporter)
}

func TestTestP01(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit1222.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Foo Su2222ite", []Reporter{junitReporter})
	//RunSpecsWithDefaultAndCustomReporters(t *testing.T, description string, reporters []Reporter)
}
