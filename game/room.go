package game

import (
	"funygame/core"
	"funygame/pb"
	"funygame/utils"
	"github.com/golang/protobuf/proto"
	"sync/atomic"
)

var roomIdGen = int64(0)

func nextRoomId() int64 {

	atomic.AddInt64(&roomIdGen, 1)

	return roomIdGen
}

type playerStatus struct {
	blood  int;
	player *Player
}

// 房间管理，保存玩家所在的房间
type RoomManager struct {
	playerRoom map[int64]*Room // 玩家所在房间

	curRoom *Room
}

func (rm *RoomManager) FindRoom(uid int64) *Room {
	if v, ok := rm.playerRoom[uid]; ok {
		return v
	}

	if rm.curRoom.hasPlayer(uid) {
		return rm.curRoom
	}

	return nil
}

func CreateRoomManager() *RoomManager {
	r := &RoomManager{
		curRoom: CreateRoom(),
	}

	return r
}

///////////////////////////////////////////////////// Room
// 一个游戏房间
type Room struct {
	status int // 0 = 初始化，1=游戏中，2=已结束

	RoomId int64

	playerStatus [10][10]*playerStatus

	// 空位
	pos []int

	// 当前空余位置
	posIndex int

	playerIndexMap map[int64]int

	MsgChan chan proto.Message
}

func CreateRoom() *Room {
	r := &Room{
		MsgChan: make(chan proto.Message),
	}
	r.RoomId = nextRoomId()
	r.pos = make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		r.pos = append(r.pos, i)
	}

	utils.Shuffle(r.pos)
	go r.run()
	return r
}

func IndexToXY(index int) [2]int {
	var a [2]int
	a[0] = index % 10
	a[1] = index / 10

	return a
}

func XYToIndex(xy [2]int) int {
	return xy[0] + xy[1]*10
}

func (r *Room) SendMsg(msg proto.Message) {
	r.MsgChan <- msg
}
func (r *Room) run() {
	select {
	case v, ok := <-r.MsgChan:
		{
			core.Debug("房间收到消息:s%v,%v", v, ok)
		}
	}
}
func (r *Room) hasPlayer(uid int64) (ok bool) {
	_, ok = r.playerIndexMap[uid]
	return
}

// 进入房间选择一个位置
func (r *Room) enterRoom(player *Player) (index int) {
	if r.status == 0 {
		index = r.posIndex
		r.posIndex++

		xy := IndexToXY(index)
		r.playerStatus[xy[0]][xy[1]] = createStatus(player)

		if r.posIndex == 100 {
			r.status = 1
			r.pushStart()
		}

		r.playerIndexMap[player.id] = index

	} else {
		index = -1
	}
	return
}

func (r *Room) exitRoom(player *Player) {

}

// 推送开始信息
func (r *Room) pushStart() {

	msg := &pb.StartGamePush_30001{
		RoomId: r.RoomId,
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			p := r.playerStatus[i][j]
			p.player.SendMsg(msg)
		}
	}
}

func createStatus(player *Player) *playerStatus {
	p := &playerStatus{blood: 2000, player: player}
	return p
}
