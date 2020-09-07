package game

import (
	"fmt"
	"squrefight/core"
)

func EnterGame(request interface{}, request2 *core.Request) interface{} {
	fmt.Println("Enter game", request)
	req := request.(*EnterGameMsg)

	GameVal.EnterGame(*req, request2)
	v := EnterGameSend{}
	v.T = req.T
	v.P = req.P
	p := GameVal.AddrMap[request2.RemoteAddr]

	p.id = req.P
	GameVal.IdMap[p.id] = p
	core.Debug("req.p:%v", req.P)
	GameVal.BroadcastMsg(v, req.P)

	for _, v := range GameVal.IdMap {

		if v.id != req.P {
			v1 := EnterGameSend{}
			v1.T = EnterGameProcess
			v1.P = v.id
			p.SendMsg(v1)
		}
	}

	return nil
}
