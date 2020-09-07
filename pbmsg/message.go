package pbmsg

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"funygame/login"
	"funygame/pb"
)

var MessageProcessMap = make(map[int32]func(*pb.Message) proto.Message)

func InitHandler() {
	fmt.Println("init handler")
	// 用户登陆
	LoginRequest(login.Login)
}
