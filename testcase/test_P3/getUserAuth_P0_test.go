package test_P3

import (
	"ccpAPITest/CCPClient/core"
	authData "ccpAPITest/data/AuthData"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

//var logger = core.Sugar

func TestP0(t *testing.T) { TestingT(t) } //继承testing的方法，可以直接使用go test命令运行

type MySuiteP0 struct{} //创建测试套件结构体

var _ = Suite(&MySuiteP0{})

func (s *MySuiteP0) TestUserAuthP1Success(c *C) {
	//logger.Info("this is test")
	conn, err := core.DialTimeout("139.199.128.158:8085", 2*time.Second)
	if err != nil {
		//logger.Infof("make session erros: %s", err)
		return
	}
	client := core.NewCCPClient(conn)
	in := authData.AuthSendData()
	out, err := client.GetAuth(in)
	if err != nil {
		//logger.Infof("client GetAuth err: %s", err)
	}
	//logger.Infof("client GetAuth success, message is: %v", out)
	c.Check(out, NotNil)
}
