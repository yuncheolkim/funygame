package game

import (
	"funygame/core"
	"github.com/golang/protobuf/proto"
	"sync"
)

type Player struct {
	id   int64
	conn *core.Connection
	data *PlayerData
	robot bool

	RoomId int32 // 所在房价id 如果不在房间 则为-1

	msgChan chan proto.Message

	mu      sync.Mutex

}

type PlayerData struct {
	Hp int
}

func (p *Player) SendMsg(i proto.Message) {
	p.mu.Lock()
	p.conn.WritePb(i)
	p.conn.Flush()
	p.mu.Unlock()
}
