package game

import "funygame/core"

type Player struct {
	id   int32
	conn *core.Connection
	data *PlayerData
	robot bool

	RoomId int32 // 所在房价id 如果不在房间 则为-1
}

type PlayerData struct {
	X  float64
	Y  float64
	Hp int
}

func (p *Player) SetId(i int32) {
	p.id = i
}

func (p *Player) SendMsg(i interface{}) {
	p.conn.Flush()
}
