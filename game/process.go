package game

import "squrefight/core"

const MoveProcess = 1
const EnterGameProcess = 2

type Process struct {
	Msg    interface{}
	Action func(interface{}, *core.Request) interface{}
}

func InitProcess() {
	GameVal.ProcessMap[MoveProcess] = Process{
		Msg:    &MoveMsg{},
		Action: MoveAction,
	}
	GameVal.ProcessMap[EnterGameProcess] = Process{
		Msg:    &EnterGameMsg{},
		Action: EnterGame,
	}
}
