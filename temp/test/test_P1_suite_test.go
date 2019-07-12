package main_test

import (
	"github.com/onsi/ginkgo/reporters"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTestP1(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "TestP1 Suite")
	//RegisterFailHandler(Fail)
	//RunSpecs(t, "TestP0 Suite")
	//junitReporter := reporters.NewJUnitReporter("junit12.html")
	junitReporter := reporters.NewJUnitReporter("j.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Foo Suite Test P1", []Reporter{junitReporter})
}
