package GetAuthSuccess_test

import (
	"ccpAPITest/CCPClient/core"
	"ccpAPITest/data/AuthData"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"time"
	//. "UNKNOWN_PACKAGE_PATH"
)

var _ = Describe("GetAuthSuccessP1", func() {
	Describe("Get Auth SuccessP1 No error", func() {
		Context("When get auth success", func() {
			conn, err := core.DialTimeout("139.199.128.158:8085", 2*time.Second)
			if err != nil {
				//logger.Infof("make session erros: %s", err)
				return
			}
			client := core.NewCCPClient(conn)
			in := AuthData.AuthSendData()
			_, err = client.GetAuth(in)
			It("should not err", func() {
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			})
		})
	})
	Describe("Get Auth SuccessP1 Out is nil", func() {
		Context("When get auth success", func() {
			conn, err := core.DialTimeout("139.199.128.158:8085", 2*time.Second)
			if err != nil {
				//logger.Infof("make session erros: %s", err)
				return
			}
			client := core.NewCCPClient(conn)
			in := AuthData.AuthSendData()
			out, err := client.GetAuth(in)
			It("out should equal nil", func() {
				gomega.Expect(out).To(gomega.Equal(nil))
			})
		})
	})
})
