package main

import (
	"bufio"
	stProto "ccpAPITest/temp/testpb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"time"
)

func main() {
	strIP := "localhost:6600"
	var conn net.Conn
	var err error

	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("Connect,", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect")
	}
	fmt.Println("connect", strIP, "success")
	defer conn.Close()

	cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		stSend := &stProto.UserInfo{
			Message: sender.Text(),
			Length:  *proto.Int(len(sender.Text())),
			Cnt:     *proto.Int(),
		}
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}
		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
	}
}
