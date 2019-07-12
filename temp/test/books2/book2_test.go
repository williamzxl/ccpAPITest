package books_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	_ "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
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
			It("Should be a novel", func() {
				//Expect().T
				//Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
				assert.Equal(GinkgoT(), longBook.CategoryByLength2(), "NOVEL")
			})

			//It("Should post tho the channel,eventually", func(done Done) {
			//	c := make(chan string, 0)
			//	go DoSomething(c)
			//	Expect(<-c).To(ContainSubstring("Done!"))
			//	close(done)
			//}, 0.2)
		})
	})
})
