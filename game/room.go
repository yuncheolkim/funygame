package game

import (
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

type Room struct {
	status int // 0 = 初始化，1=游戏中，2=已结束

	RoomId int64

	playerStatus [10][10]*playerStatus

	pos []int

	// 当前空余位置
	posIndex int
}

func CreateRoom() *Room {
	r := &Room{}
	r.RoomId = nextRoomId()
	r.pos = make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		r.pos = append(r.pos, i)
	}

	utils.Shuffle(r.pos)

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

// 进入房间选择一个位置
func (r *Room) EnterRoom(player *Player) (index int) {
	if r.status == 0 {
		index = r.posIndex
		r.posIndex++

		xy := IndexToXY(index)
		r.playerStatus[xy[0]][xy[1]] = createStatus(player)

		if r.posIndex == 100 {
			r.status = 1
			r.pushStart()
		}

	} else {
		index = -1
	}
	return
}

func (r *Room) ExitRoom(player *Player) {

}

// 推送开始信息
func (r *Room) pushStart() {

	msg := pb.StartGamePush_30001{
		RoomId: r.RoomId,
	}

	for i:= 0;i<10;i++{
		for j := 0 ;j< 10; j++{
			p := r.playerStatus[i][j]
			p.player.SendMsg(msg)
		}
	}
}

func createStatus(player *Player) *playerStatus {
	p := &playerStatus{blood: 2000,player:player}
	return p
}


///////////////// action
// 开始匹配
func MatchAction(msg interface{}, player *Player) proto.Message {


	return nil
}