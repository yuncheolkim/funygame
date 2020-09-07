package game

type Msg struct {
	T int32 // 消息类型
	P int32 // 用户id
	S int64 // 消息序列
}

type MsgRecv struct {
	Msg
}

// 移动
type MoveMsg struct {
	MsgRecv
	X float64
	Y float64
}

type EnterGameMsg struct {
	MsgRecv
}

type EnterGameSend struct {
	Msg
}
