package books_test

import (
	. "github.com/onsi/ginkgo"
	_ "github.com/smartystreets/goconvey/convey"
	//"github.com/stretchr/testify/assert"
	. "github.com/onsi/gomega"
	//. "github.com/onsi/gomega"
	//. "UNKNOWN_PACKAGE_PATH"
)

type Book2 struct {
	Title  string
	Author string
	Pages  int
}

func (book Book) CategoryByLength2() string {
	if book.Pages > 1000 {
		return "NOVEL"
	}
	return "SMALL"
}

var _ = Describe("Book", func() {
	var (
		longBook Book
		//shortBook Book
	)
	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}
		//shortBook =Book{
		//	Title:  "Fox In Socks",
		//	Author: "Dr. Seuss",
		//	Pages:  24,
		//}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("Books2 Should be a novel", func() {
				//Expect().T
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
				//assert.Equal(GinkgoT(), longBook.CategoryByLength2(), "NOVEL")
			})
		})
	})
})
