package AuthData

import (
	"ccpAPITest/CCPClient/protobuf/auth"
	"crypto/md5"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"
	"time"
)

func AuthSendData() *auth.UserAuthInner {
	timeStamp := time.Now().Format("20060102150405")
	data := []byte("ff8080815dbc080c015dbc9d7cd40003" + "963" + timeStamp + "7f4fa6d320ab49739183af1d498adb6b")
	md5 := fmt.Sprintf("%x", md5.Sum(data))
	in := &auth.UserAuthInner{
		AuthType:   proto.Uint32(1),
		AppId:      proto.String("ff8080815dbc080c015dbc9d7cd40003"),
		UserName:   proto.String("963"),
		Device:     proto.Uint32(20),
		Timestamp:  proto.String(timeStamp),
		Sig:        proto.String(strings.ToUpper(md5)),
		SdkVersion: proto.String("5.4.2r"),
		Imei:       proto.String("990009263657473"),
		Mode:       proto.Uint32(2),
		Network:    proto.Uint32(1),
	}
	return in
}
