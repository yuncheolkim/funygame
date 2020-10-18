package game

import (
	"fmt"
	"funygame/core"
	"funygame/pb"
	"github.com/golang/protobuf/proto"
	"sync"
)

type Player struct {
	id    int64
	conn  *core.Connection
	data  *PlayerData
	robot bool

	RoomId int32 // 所在房价id 如果不在房间 则为-1

	msgChan chan proto.Message

	mu sync.Mutex
}

type PlayerData struct {
	Hp  int32
	Def int32
}

func (p *Player) GetId() int64 {
	return p.id
}

func (p *Player) SendMsg(i proto.Message, msgNo int32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Printf("[%v]发送消息，消息号:%v\n",p.id, msgNo)
	bytes, _ := proto.Marshal(i)
	m := pb.Message{
		Seq:          0,
		MsgNo:        msgNo,
		BroadcastUid: nil,
		Body:         bytes,
		Uid:          0,
	}
	p.conn.WritePb(&m)
	p.conn.Flush()

}

// 受到攻击
func (p *Player) attacked(damage int32) int32 {
	if damage <= 0{
		return 0
	}
	if p.data.Def <= 0 {
		return damage
	}

	if p.data.Def >= damage {
		p.data.Def -= damage
		return 0
	}

	i := damage - p.data.Def

	p.data.Hp -= i;
	return i
}

func (p *Player) addDef(v int32) {
	p.data.Def += v
	if p.data.Def > 500 {
		p.data.Def = 500
	}
}

func (p *Player) addHp(i int32) int32 {
	old := p.data.Hp
	p.data.Hp += i
	if p.data.Hp > 2000 {
		p.data.Hp = 2000
	}

	return p.data.Hp - old
}
