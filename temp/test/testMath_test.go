package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"

	//. "UNKNOWN_PACKAGE_PATH"
	. "github.com/onsi/ginkgo/extensions/table"
)

var _ = Describe("TestMath", func() {
	DescribeTable("the > inequality",
		func(x int, y int, expected bool) {
			Expect(x > y).To(Equal(expected))
		},
		Entry("x > y", 1, 0, true),
		Entry("x == y", 1, 1, false),
		Entry("x < y", 0, 1, false),
	)
})

var _ = Describe("Substring matching", func() {
	type SubstringCase struct {
		String    string
		Substring string
		Count     int
	}

	DescribeTable("Countting substring matches",
		func(c SubstringCase) {
			Ω(strings.Count(c.String, c.Substring)).Should(BeNumerically("==", c.Count))
		},
		Entry("With no matching substring", SubstringCase{
			String:    "The sixth shei",
			Substring: "W",
			Count:     0,
		}),
		Entry("1", SubstringCase{
			String:    "fsfsf ddd",
			Substring: "ddd",
			Count:     12,
		}),
		Entry("2", SubstringCase{
			String:    "fsfsf ddd ddd",
			Substring: "ddd",
			Count:     2,
		}),
	)
})
