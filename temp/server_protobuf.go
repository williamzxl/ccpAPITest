package main

import (
	stProto "ccpAPITest/temp/testpb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "local host：6600")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("new connect", conn.RemoteAddr())
		go readMessage(conn)
	}
}

func readMessage(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		cnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		stReceive := &stProto.UserInfo{}
		pData := buf[:cnt]
		err = proto.Unmarshal(pData, stReceive)
		if err != nil {
			panic(err)
		}
		fmt.Println("receive", conn.RemoteAddr(), stReceive)
		if stReceive.Message == "stop" {
			os.Exit(1)
		}
	}

}
