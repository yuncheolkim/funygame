package game

import (
	"funygame/pb"
	"github.com/golang/protobuf/proto"
)

const MoveProcess = 1
const EnterGameProcess = 2

type Action func(msg interface{}, player *Player) proto.Message

type Process struct {
	Msg    func() proto.Message
	Action Action
}

func InitProcess() {
	GameVal.ProcessMap[MoveProcess] = Process{
		Action: MatchAction,
		Msg: func() proto.Message {
			return &pb.StartMatchReq_10001{}
		},
	}
}
