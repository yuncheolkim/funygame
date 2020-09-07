package game

import "funygame/core"

type Player struct {
	id   int32
	conn *core.Connection

	X  float64
	Y  float64
	Hp int
}

func (p *Player) SetId(i int32) {

	p.id = i
}

func (p *Player) SendMsg(i interface{}) {

	p.conn.WriteJson(i)
	p.conn.Flush()

}
