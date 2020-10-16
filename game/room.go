package game

import (
	"funygame/pb"
	"funygame/utils"
	"github.com/golang/protobuf/proto"
	"sync"
	"sync/atomic"
)

var roomIdGen = int64(0)

func nextRoomId() int64 {

	atomic.AddInt64(&roomIdGen, 1)

	return roomIdGen
}

type playerStatus struct {
	player *Player
}

// 房间管理，保存玩家所在的房间
type RoomManager struct {
	mu         sync.Mutex
	playerRoom map[int64]*Room // 玩家所在房间
	curRoom    *Room
}

func (rm *RoomManager) FindRoom(p *Player) *Room {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if v, ok := rm.playerRoom[p.id]; ok {
		return v
	}

	if rm.curRoom.hasPlayer(p.id) {
		return rm.curRoom
	}

	return nil
}

func (rm *RoomManager) JoinRoom(player *Player) *Room {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if v, ok := rm.playerRoom[player.id]; ok {
		return v
	}

	if !rm.curRoom.hasPlayer(player.id) {
		rm.curRoom.enterRoom(player)

		rm.playerRoom[player.id] = rm.curRoom
		if rm.curRoom.isStart() {
			rm.curRoom = CreateRoom()
		}
	}

	return rm.playerRoom[player.id]
}

func CreateRoomManager() *RoomManager {
	r := &RoomManager{
		curRoom:    CreateRoom(),
		playerRoom: make(map[int64]*Room),
	}

	return r
}

///////////////////////////////////////////////////// Room
// 一个游戏房间
type Room struct {
	status int // 0 = 初始化，1=游戏中，2=已结束

	RoomId int64

	playerStatus [100]*playerStatus

	// 空位
	pos []int

	// 当前空余位置
	posIndex int

	playerIndexMap map[int64]int

	mu sync.Mutex
}

func CreateRoom() *Room {
	r := &Room{
		playerIndexMap: make(map[int64]int),
	}
	r.RoomId = nextRoomId()
	r.pos = make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		r.pos = append(r.pos, i)
	}

	utils.Shuffle(r.pos)
	return r
}


func (r *Room) hasPlayer(uid int64) (ok bool) {
	_, ok = r.playerIndexMap[uid]
	return
}

// 玩家进入房间
func (r *Room) enterRoom(player *Player) (index int) {
	if r.status == 0 {

		if v, ok := r.playerIndexMap[player.id]; ok {
			return v
		}
		index = r.posIndex
		r.posIndex++

		r.playerStatus[index] = createStatus(player)

		r.broadcastPlayerJoin(player)
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

func (r *Room) isStart() bool {
	return r.status == 1
}

// 推送开始信息
func (r *Room) pushStart() {

	msg := &pb.StartGamePush_30001{
		RoomId: r.RoomId,
	}

	for i := 0; i < 100; i++ {
		p := r.playerStatus[i]
		p.player.SendMsg(msg, 30001)
	}
}

func (r *Room) GetIndex(uid int64) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.playerIndexMap[uid]
}

func (r *Room) broadcastPlayerJoin(player *Player) {

	msg := pb.UserEnterPush_30002{
		Index: int32(r.GetIndex(player.id)),
	}
	for i := 0; i < 100; i++ {
		s := r.playerStatus[i]
		if s != nil && !s.player.robot && s.player.id != player.id {
			s.player.SendMsg(&msg, 30002)
		}
	}
}

func (r *Room) broadcast(m proto.Message, msgNo int32) {

	for i := 0; i < 100; i++ {
		s := r.playerStatus[i]
		if s != nil && !s.player.robot {
			s.player.SendMsg(m, msgNo)
		}
	}
}

// 攻击敌人
func (r *Room) attack(player *Player, index int32, damage int32) {
	r.mu.Lock()
	r.mu.Unlock()

	p := r.playerStatus[index].player

	attacked := p.attacked(damage)
	if attacked > 0 { // 掉血，给所有人推送
		m := pb.BloodChangePush_30003{
			Index: index,
			Num:   attacked * (-1),
		}

		r.broadcast(&m, 30003)
	} else {
		// 给本人发送消息，减少护盾
		m := pb.BloodChangePush_30003{
			Index: index,
			Num:   attacked * (-1),
		}
		r.playerStatus[index].player.SendMsg(&m, 30003)
	}
}

func (r *Room) addDef(player *Player, i int32) {
	r.mu.Lock()
	r.mu.Unlock()
	player.addDef(i)
}

func (r *Room) addHp(player *Player, i int32) {
	r.mu.Lock()
	r.mu.Unlock()

	added := player.addHp(i)
	if added > 0 {
		m := pb.BloodChangePush_30003{
			Index: int32(r.playerIndexMap[player.id]),
			Num:   added,
		}
		r.broadcast(&m, 30003)
	}

}

func createStatus(player *Player) *playerStatus {
	p := &playerStatus{player: player}
	return p
}