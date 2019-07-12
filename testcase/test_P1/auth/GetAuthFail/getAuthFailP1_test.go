package GetAuthFail_test

import (
	"ccpAPITest/CCPClient/core"
	"ccpAPITest/data/AuthData"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"time"
	//. "UNKNOWN_PACKAGE_PATH"
)

var _ = Describe("GetAuthFail P1", func() {
	Describe("Get Auth Fail P1 error", func() {
		Context("When get auth fail", func() {
			conn, err := core.DialTimeout("139.199.128.158:8085", 2*time.Second)
			//if err != nil {
			//	//logger.Infof("make session erros: %s", err)
			//	return
			//}
			client := core.NewCCPClient(conn)
			in := AuthData.AuthSendData()
			_, err = client.GetAuth(in)
			It("should be err", func() {
				gomega.Expect(err).To(gomega.HaveOccurred())
			})
		})
	})
	Describe("Get Auth P1 Fail Out is 0", func() {
		Context("When get auth fail", func() {
			conn, _ := core.DialTimeout("139.199.128.158:8085", 2*time.Second)
			//if err != nil {
			//	//logger.Infof("make session erros: %s", err)
			//	return
			//}
			client := core.NewCCPClient(conn)
			in := AuthData.AuthSendData()
			out, _ := client.GetAuth(in)
			It("out should equal 0", func() {
				gomega.Expect(out).To(gomega.Equal(0))
			})
		})
	})
})
