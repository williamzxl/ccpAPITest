package test_P3

import (
	"ccpAPITest/CCPClient/core"
	authData "ccpAPITest/data/AuthData"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

var logger = core.Sugar

func TestP00(t *testing.T) { TestingT(t) } //继承testing的方法，可以直接使用go test命令运行

type MySuite struct{} //创建测试套件结构体

var _ = Suite(&MySuite{})

func (s *MySuite) TestUserAuthSuccess(c *C) {
	conn, err := core.DialTimeout("139.199.128.158:8085", 2*time.Second)
	if err != nil {
		return
	}
	client := core.NewCCPClient(conn)
	in := authData.AuthSendData()
	out, err := client.GetAuth(in)
	if err != nil {
	}
	c.Check(out, NotNil)
	//convey.Convey("Check OUT", c, func() {
	//	convey.So(out, convey.ShouldEqual, nil)
	//})
}

func (s *MySuite) TestUserAuthFail(c *C) {
	conn, err := core.DialTimeout("139.199.128.158:8085", 2*time.Second)
	if err != nil {
		return
	}
	client := core.NewCCPClient(conn)
	in := authData.AuthSendData()
	out, err := client.GetAuth(in)
	if err != nil {
	}
	c.Check(out, Equals, 0)
	//convey.Convey("Check OUT", c, func() {
	//	convey.So(out, convey.ShouldEqual, "")
	//})
}
