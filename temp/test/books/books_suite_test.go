package books_test

import (
	"testing"
	//"db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "github.com/stretchr/testify/assert"
)

//var dbRunner *db.Runner
//var dbClient *db.Client
//
//var _ = BeforeSuite(func() {
//	dbRunner = db.NewRunner()
//})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite Book1")
}
