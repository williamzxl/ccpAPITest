package core

import (
	"log"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	//"github.com/smartystreets/goconvey"
)

var (
	ReceiveTimeoutError = errors.New("ReceiveTimeoutError")
	StopWaitError       = errors.New("session closed, stop waiting")
)

type codecPack struct {
	MsgType uint32
	Data    []byte
}

func (c *CCPClient) Invoke(method string, req, reply interface{}) error {
	return invoke(method, req, reply, c)
}

// TODO：增加输入函数参数 opts ...CallOption 的处理
func invoke(method string, req, reply interface{}, client *CCPClient) error {
	protoType, err := method2ProtoType(method)
	if err != nil {
		logger.Debugf("method not found: %v", method)
		return errors.Wrap(err, "invoke method2ProtoType err")
	}
	sendData, err := proto.Marshal(req.(proto.Message))
	if err != nil {
		logger.Debugf("message Marshal error: %s", err)
		return errors.Wrap(err, "invoke Marshal errr")
	}
	packet := codecPack{
		MsgType: protoType,
		Data:    sendData,
	}
	if err := client.session.Send(packet); err != nil {
		logger.Debugf("session send error: %v", packet)
		return errors.Wrap(err, "invoke session send err")
	}

	receiveData, err := waitRespone(protoType, client)
	if err != nil {
		return errors.Wrap(err, "invoke waitRespone err")
	}
	if err := proto.Unmarshal(receiveData.([]byte), reply.(proto.Message)); err != nil {
		logger.Debugf("invoke Unmarshal error: %s", err)
		return errors.Wrap(err, "invoke Unmarshal data err")
	}
	logger.Debugf("session receive message: %v", reply)
	return errors.Wrap(err, "invoke err")
}

func waitRespone(msgType uint32, client *CCPClient) (interface{}, error) {
	// checkTime := time.NewTicker(time.Millisecond * 100)
	timeout := time.NewTicker(time.Second * 6)
	// defer checkTime.Stop()
	defer timeout.Stop()
	dataChan := getMapChan(&client.recvDataMap, msgType)
	for {
		select {
		case data := <-dataChan:
			logger.Info("receive data from dataChan!!!")
			// data, isPresence := client.recvDataMap.Load(uint32(1))
			return data, nil
		case <-client.sessionStopChan:
			return nil, StopWaitError
		case <-timeout.C:
			logger.Info("timout receive message")
			return nil, ReceiveTimeoutError
		}
	}
}

func getMapChan(syncMap *sync.Map, mapKey uint32) chan interface{} {
	if dataChan, isPresence := syncMap.Load(mapKey); isPresence {
		return dataChan.(chan interface{})
	}
	dataChan := make(chan interface{})
	syncMap.Store(mapKey, dataChan)
	return dataChan
}

func (c *CCPClient) receiveLoop() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		select {
		case <-c.sessionStopChan:
			logger.Debug("close chan closed(), receiveLoop stop")
			return
		case <-ticker.C:
			r, err := c.session.Receive()
			if err == HEAD_EOF_ERRR {
				continue
			}
			if err != nil {
				log.Print("receiveLoop ERROR:", err)
				continue
			}
			// c.recvDataMap.Store(r.(codecPack).MsgType, r)
			dataChan := getMapChan(&c.recvDataMap, r.(codecPack).MsgType)
			dataChan <- r.(codecPack).Data
			// c.recvDataMap.Range(func(k, v interface{}) bool {
			// logger.Debug(k, v)
			// return true
			// })
		}
	}
}
