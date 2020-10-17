package game

import (
	"funygame/core"
	"funygame/pb"
	"github.com/golang/protobuf/proto"
)

///////////////// action
// 开始匹配
func MatchAction(msg proto.Message, player *Player) proto.Message {

	//req := msg.(*pb.StartMatchReq_10001)

	room := GameVal.RoomManager.FindRoom(player)

	if room == nil {
		room = GameVal.RoomManager.JoinRoom(player)
	}

	res := &pb.StartMatchRes_10001{
		Index: int32(room.GetIndex(player.GetId())),
	}
	return res
}

// 攻击
func AttackAction(msg proto.Message, player *Player) proto.Message {

	room := GameVal.RoomManager.FindRoom(player)
	if room != nil {
		m := msg.(*pb.AttackTell_20001)
		core.Logf("攻击")
		room.attack(player, m.Index, m.Num)
	}

	return nil
}

// 回血
func CureAction(msg proto.Message, player *Player)proto.Message{

	room := GameVal.RoomManager.FindRoom(player)
	if room != nil {
		m := msg.(*pb.CureTell_20002)
		core.Logf("加血")
		room.addHp(player, m.Num)
	}
	return nil
}

// 防御
func DefAction(msg proto.Message, player *Player)proto.Message {

	room := GameVal.RoomManager.FindRoom(player)
	if room != nil {
		m := msg.(*pb.DefTell_20003)
		core.Logf("加防")
		room.addDef(player, m.Num)
	}
	return nil
}
