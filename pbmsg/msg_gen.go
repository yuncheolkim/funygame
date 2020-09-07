package pbmsg

import (
	"github.com/golang/protobuf/proto"
	"funygame/pb"
)

func LoginRequest(process func(*pb.LoginRequest) proto.Message) {
	MessageProcessMap[11] = func(message *pb.Message) proto.Message {
		return process(message.Content.(*pb.Message_LoginRes).LoginRes)
	}
}
