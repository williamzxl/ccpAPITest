package core

import (
	"bytes"
	"io"

	"github.com/funny/link"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	//"CCPClient/protobuf/comm"
	"ccpAPITest/CCPClient/protobuf/comm"
)

var (
	HEAD_EOF_ERRR = errors.New("HEAD_EOF_ERRR")
)

type CCPProtocol struct {
}

func newCCPProtocol() *CCPProtocol {
	return &CCPProtocol{}
}

func (c *CCPProtocol) NewCodec(rw io.ReadWriter) (link.Codec, error) {
	codec := &CCPCodec{
		p:  c,
		rw: rw,
	}
	return codec, nil
}

type CCPCodec struct {
	p       *CCPProtocol
	rw      io.ReadWriter
	recvBuf bytes.Reader // TODO:收包缓存还没有用起来
	sendBuf bytes.Buffer
}

func (c *CCPCodec) Send(msg interface{}) error {
	lite := &comm.MsgLiteInner{}
	lite.ProtoType = proto.Uint32(msg.(codecPack).MsgType)
	lite.ProtoClientNo = proto.Uint32(1)
	sendData := msg.(codecPack).Data
	lite.ProtoData = varintEncode(sendData)
	out, err := proto.Marshal(lite)
	if err != nil {
		logger.Debug("MsgLiteInner Marshal err %s", err)
		return errors.Wrap(err, "MsgLiteInner Marshal err")
	}
	c.sendBuf.Reset()
	c.sendBuf.Write(varintEncode(out))
	logger.Debug("session send message:", lite)
	_, err = c.rw.Write(c.sendBuf.Bytes())
	return errors.Wrap(err, "session send errwr")
}

func varintEncode(data []byte) []byte {
	lenBuf := proto.EncodeVarint(uint64(len(data)))
	return append(lenBuf, data...)
}

func varintDecode(data []byte) []byte {
	_, headerLen := proto.DecodeVarint(data)
	logger.Debug("message head len:", headerLen)
	return data[headerLen:]
}

func (c *CCPCodec) Receive() (interface{}, error) {
	// read head buff
	headBuf := make([]byte, 1)
	_, err := io.ReadFull(c.rw, headBuf)
	if err == io.EOF {
		return nil, HEAD_EOF_ERRR
	}
	if err != nil {
		logger.Debug("head read err:", err)
		return nil, err
	}
	msgLen, _ := proto.DecodeVarint(headBuf)
	logger.Debugf("msgLen = %v", msgLen)
	// read head buff
	msgBuf := make([]byte, msgLen)
	_, err = io.ReadFull(c.rw, msgBuf)
	if err != nil {
		logger.Debug("body read err:", err)
		return nil, errors.Wrap(err, "body read err")
	}
	logger.Debugf("body is %v", msgBuf)

	// Unmarshal message body
	lite := &comm.MsgLiteInner{}
	err = proto.Unmarshal(msgBuf, lite)
	if err != nil {
		logger.Debug("Unamrshal error")
		return nil, err
	}
	logger.Debug("receive msg is:", lite)
	if *lite.ProtoErrorCode > 0 {
		logger.Infow(
			"Request rejection",
			"ProtoType", int(*lite.ProtoType),
			"ProtoErrorCode", int(*lite.ProtoErrorCode),
		)
		// err = errors.New("Request rejection")
	}
	out := codecPack{
		MsgType: *lite.ProtoType,
		// Data:    varintDecode(lite.ProtoData),
		Data: lite.ProtoData,
	}
	return out, errors.Wrap(err, "session reveive err")
}

func (c *CCPCodec) Close() error {
	logger.Debug("session closed")
	return nil
}
