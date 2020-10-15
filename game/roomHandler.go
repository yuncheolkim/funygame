package game

import (
	"funygame/pb"
	"github.com/golang/protobuf/proto"
)

///////////////// action
// 开始匹配
func MatchAction(msg interface{}, player *Player) proto.Message {

	req := msg.(*pb.StartMatchReq_10001)

	room :=	GameVal.RoomManager.FindRoom(player)

	if room == nil{
		GameVal.RoomManager.JoinRoom(player)
	}

	print(req)
	return &pb.Success{Ok: "ok"}
}
