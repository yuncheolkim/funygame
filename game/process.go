package game

import (
	"funygame/pb"
	"github.com/golang/protobuf/proto"
)

type Action func(msg interface{}, player *Player) proto.Message

type Process struct {
	Msg    func() proto.Message
	Action Action
}

func InitProcess() {
	GameVal.ProcessMap[10001] = Process{
		Action: MatchAction,
		Msg: func() proto.Message {
			return &pb.StartMatchReq_10001{}
		},
	}
}
