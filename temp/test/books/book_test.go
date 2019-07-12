package books_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "github.com/smartystreets/goconvey/convey"
	"strings"
	//. "github.com/onsi/gomega"
	//. "UNKNOWN_PACKAGE_PATH"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) AuthoLastName() string {
	var temp string
	temp = b.Author
	return strings.Split(temp, " ")[1]
}

func (book Book) CategoryByLength() string {
	if book.Pages > 1000 {
		return "NOVEL"
	}
	if book.Pages < 100 {
		return "SHORT"
	}
	return "SMALL"
}

//var _ = Describe("Book", func() {
//	 var (
//	 	longBook Book
//	 	shortBook Book
//	 )
//	 BeforeEach(func() {
//	 	longBook = Book{
//			Title:  "Les Miserables",
//			Author: "Victor Hugo",
//			Pages:  1488,
//		}
//	 	shortBook =Book{
//			Title:  "Fox In Socks",
//			Author: "Dr. Seuss",
//			Pages:  24,
//		}
//	 })
//
//	 Describe("Categorizing book length", func() {
//		 Context("With more than 300 pages", func() {
//			 It("Books Should be a novel", func() {
//				  //Expect().T
//			 	//Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
//			 	assert.Equal(GinkgoT(), longBook.CategoryByLength(), "NOVEL")
//			 })
//		 })
//	 })
//
//	 Describe("Categorizing book length2", func() {
//		Context("Books With less than 100 pages", func() {
//			It("Should be a short", func() {
//				//Expect().T
//				//Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
//				assert.Equal(GinkgoT(), shortBook.CategoryByLength(), "SHORT")
//			})
//		})
//	})
//})

//type NewBookFromJSON struct {
//	title string `json:"title"`
//	author string `json:"author"`
//	pages int `json:"pages"`
//}

//var _ =Describe("Test Books", func() {
//	var book Book
//	BeforeEach(func() {
//		book = NewBookFromJSON(`{
//            "title":"Les Miserables",
//            "author":"Victor Hugo",
//            "pages":1488
//        }`)
//	})
//	It("Can be loaded from json", func() {
//		Expect(book.Title).To(Equal("Les Miserables"))
//		Expect(book.Author).To(Equal("Victor Hugo"))
//		Expect(book.Pages).To(Equal(1488))
//	})
//})

func NewBookFromJSON(s string) (b Book, err error) {
	//var b Book
	err = json.Unmarshal([]byte(string(s)), &b)
	return b, err
}

var _ = Describe("Book", func() {
	var (
		book Book
		err  error
	)
	BeforeEach(func() {
		book, err = NewBookFromJSON(`{
			"title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
		}`)
	})
	Describe("Loading form JSON", func() {
		Context("When the JSON parses successfully", func() {
			It("Should populate the fields correctly", func() {
				By("Test1: Title")
				Expect(book.Title).To(Equal("Les Miserables"))
				By("Test2:Author")
				Expect(book.Author).To(Equal("Victor Hugo"))
				By("Test3:Pages")
				Expect(book.Pages).To(Equal(14088))
			})
			It("Should not error", func() {
				Expect(nil).NotTo(HaveOccurred())
			})
		})

		Context("When the JSON fails to parse", func() {
			BeforeEach(func() {
				book, err = NewBookFromJSON(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":oo1488oops
                }`)
			})
			It("Should return the zero-value for the boook", func() {
				Expect(book).To(BeZero())
			})

			It("Shold error", func() {
				//fmt.Println("Err:", err)
				Expect(err).To(HaveOccurred())
			})
		})

		//Measure("it should do something hard efficiently", func(b Benchmarker) {
		//	runtime := b.Time("runtime", func() {
		//		output := SomethingHard()
		//		Expect(output).To(Equal(19))
		//	})
		//})
		Describe("My api client", func() {
			var client APIClient
			var fakeServer FakerServer
			var response chan APIREsponse

			BeforeEach(func() {
				resqponse = make(chan APIResponse, 1)
				fakeServer = NewFAkeServer()
				client = NewAPIClient(fakeServer)
				client.Get("/some/endpoint", response)
			})
			Describe("Failure modes", func() {
				AssertFailedBehavior := func() {
					It("Shold not include", func() {
						Ω((<-response).JSON).Should(BeZero())
					})
					It("shold not report", func() {
						Ω((<-response).JSON).Should(BeZero())
					})
				}

				Context("when", func() {
					BeforeEach(func() {
						fakeServer.Response(404)
					})
					AssertFailedBehavior()
				})
			})

		})
	})

	Describe("Extracting the author's last name", func() {
		It("Shold correctly identify and return the last name", func() {
			//fmt.Println("Focus")
			Expect(book.AuthoLastName()).To(Equal("Hugo"))
		})
	})
})
