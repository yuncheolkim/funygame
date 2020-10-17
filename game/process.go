package game

import (
	"funygame/pb"
	"github.com/golang/protobuf/proto"
)

type Action func(msg proto.Message, player *Player) proto.Message

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
	GameVal.ProcessMap[20001] = Process{
		Action: AttackAction,
		Msg: func() proto.Message {
			return &pb.AttackTell_20001{}
		},
	}
	GameVal.ProcessMap[20002] = Process{
		Action: CureAction,
		Msg: func() proto.Message {
			return &pb.CureTell_20002{}
		},
	}
	GameVal.ProcessMap[20003] = Process{
		Action: DefAction,
		Msg: func() proto.Message {
			return &pb.DefTell_20003{}
		},
	}
}
