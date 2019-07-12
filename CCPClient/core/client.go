package core

import (
	"ccpAPITest/CCPClient/protobuf/auth"
	"sync"
	"time"

	"github.com/funny/link"
	"github.com/pkg/errors"
)

type CCPClient struct {
	session         *link.Session
	recvDataMap     sync.Map
	recvStopChan    chan bool
	sessionStopChan chan bool
}

func Dial(address string) (*link.Session, error) {
	protocol := newCCPProtocol()
	//it will reutrn a link.Session, it has Close method.
	// 返回一个 link.Session 对象，会自己调用 close 方法
	return link.Dial("tcp", address, protocol, 0)
}

func DialTimeout(address string, timeout time.Duration) (*link.Session, error) {
	protocol := newCCPProtocol()
	//it will reutrn a link.Session, it has Close method.
	return link.DialTimeout("tcp", address, timeout, protocol, 0)
}
func NewCCPClient(session *link.Session) *CCPClient {
	c := &CCPClient{
		session:         session,
		recvStopChan:    make(chan bool),
		sessionStopChan: make(chan bool),
	}
	callback := func() {
		close(c.sessionStopChan)
	}
	c.session.AddCloseCallback(nil, 1, callback)
	logger.Debugf("new ccp client: %v", c)
	go c.receiveLoop()
	return c
}

// TODO：增加输入函数参数 opts ...CallOption 的处理
func (c *CCPClient) GetAuth(in *auth.UserAuthInner) (*auth.UserAuthRespInner, error) {
	out := &auth.UserAuthRespInner{}
	err := c.Invoke("GetAuth", in, out)
	if err != nil {
		return nil, errors.Wrap(err, "client GetAuth err")
	}
	return out, nil
}
