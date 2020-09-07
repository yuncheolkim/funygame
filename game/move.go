package game

import (
	"fmt"
	"funygame/core"
)

func MoveAction(request interface{}, request2 *core.Request) interface{} {
	fmt.Println("move action", request)
	req := request.(*MoveMsg)
	v := MoveMsg{}
	v.T = req.T
	v.P = req.P
	v.X = req.X
	v.Y = req.Y

	GameVal.BroadcastMsg(v, req.P)

	return nil
}
