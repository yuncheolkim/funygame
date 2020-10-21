package game

import (
	"fmt"
	"funygame/pb"
	"funygame/utils"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
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
		if v.isEnd(true) {
			return nil
		}
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
		if !v.isStart(true)&&!v.isEnd(true) {
			return v
		}
	}

	if !rm.curRoom.isInit(true){
		rm.curRoom = CreateRoom()
	}

	if !rm.curRoom.hasPlayer(player.id) {
		rm.curRoom.enterRoom(player)

		rm.playerRoom[player.id] = rm.curRoom
		if rm.curRoom.isStart(true) {
			r := CreateRoom()
			r.rm = rm
			rm.curRoom = r
		}
	}

	return rm.playerRoom[player.id]
}

func (rm *RoomManager) ExitRoom(p *Player) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if v, ok := rm.playerRoom[p.id]; ok {
		v.exitRoom(p)
		delete(rm.playerRoom,p.id)
	} else if rm.curRoom.hasPlayer(p.id) {
		rm.curRoom.exitRoom(p)
	}
}

func CreateRoomManager() *RoomManager {
	room := CreateRoom()
	r := &RoomManager{
		curRoom:    room,
		playerRoom: make(map[int64]*Room),
	}
	r.curRoom.rm = r

	return r
}

///////////////////////////////////////////////////// Room
// 一个游戏房间
type Room struct {
	rm     *RoomManager
	status int // 0 = 初始化，1=游戏中，2=已结束

	RoomId int64

	playerStatus [100]*playerStatus

	// 空位
	pos []int

	// 当前空余位置
	posIndex int

	playerIndexMap map[int64]int
	humanIndexMap  map[int64]int
	alive          []int

	mu sync.Mutex

	tick *time.Ticker
}

func CreateRoom() *Room {
	r := &Room{
		playerIndexMap: make(map[int64]int),
		humanIndexMap:  make(map[int64]int),
		tick:           time.NewTicker(time.Second),
	}
	r.RoomId = nextRoomId()
	r.pos = make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		r.pos = append(r.pos, i)
	}

	utils.Shuffle(r.pos)
	go r.addRobotRun()
	return r
}

func (r *Room) hasPlayer(uid int64) (ok bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok = r.playerIndexMap[uid]
	return
}

// 玩家进入房间
func (r *Room) enterRoom(player *Player) (index int) {

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.status == 0 {

		if v, ok := r.playerIndexMap[player.id]; ok {
			return v
		}
		index = r.pos[r.posIndex]
		r.posIndex++

		if !player.robot {
			r.humanIndexMap[player.id] = index
		}

		r.alive = append(r.alive, index)

		r.playerStatus[index] = createStatus(player)
		r.playerIndexMap[player.id] = index
		r.broadcastPlayerJoin(player)
		if r.posIndex == 100 {
			r.status = 1
			r.pushStart()
		}

	} else {
		index = -1
	}
	return
}

func (r *Room) addRobotRun() {
	for {
		r.mu.Lock()
		if r.status == 2 {
			r.tick.Stop()
			r.mu.Unlock()
			fmt.Println("End game")
			break
		}
		r.mu.Unlock()
		select {
		case <-r.tick.C:
			{
				r.mu.Lock()
				if r.isStart(false) {
					//机器人攻击
					for _, v := range r.playerStatus {
						if v != nil && v.player.robot && v.player.IsAlive() {
							r.attack(v.player, int32(r.randIndex(r.playerIndexMap[v.player.id])), 10, false)
						}
					}

					r.mu.Unlock()
				} else if len(r.humanIndexMap) > 0 {
					r.mu.Unlock()
					// 加机器人
					r.enterRoom(CreateRobot())
				} else {
					r.mu.Unlock()

				}

			}
		}
	}

}

func (r *Room) exitRoom(player *Player) {
	index := r.playerIndexMap[player.id]
	r.playerStatus[index] = nil

	r.posIndex -= 1
	if r.posIndex < 0 {
		r.posIndex = 0
	}
	// 发送消息
	msg := &pb.LeaveRoomReq_30004{
		Index: int32(index),
	}
	r.broadcast(msg, 30004)
}

func (r *Room) isStart(lock bool) bool {
	if lock {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	return r.status == 1
}

// 推送开始信息
func (r *Room) pushStart() {

	msg := &pb.StartGamePush_30001{
		RoomId: r.RoomId,
	}

	for i := 0; i < 100; i++ {
		p := r.playerStatus[i]
		if p != nil {
			p.player.SendMsg(msg, 30001)
		}
	}

}

func (r *Room) GetIndex(uid int64) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.playerIndexMap[uid]
}
func (r *Room) GetIndexInLock(uid int64) int {
	return r.playerIndexMap[uid]
}

func (r *Room) broadcastPlayerJoin(player *Player) {

	msg := pb.UserEnterPush_30002{
		Index: int32(r.playerIndexMap[player.id]),
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
func (r *Room) attack(player *Player, index int32, damage int32, lock bool) {
	if lock {
		r.mu.Lock()
		defer r.mu.Unlock()
	}
	if !r.isStart(false) {
		return
	}
	if r.playerStatus[index] == nil {
		return
	}
	p := r.playerStatus[index].player

	attacked := p.attacked(damage)
	if attacked > 0 { // 掉血，给所有人推送
		m := pb.BloodChangePush_30003{
			Index: index,
			Num:   attacked * (-1),
		}
		r.broadcast(&m, 30003)

		if p.data.Hp <= 0 {
			r.alive = utils.DeleteSlice(r.alive, r.playerIndexMap[p.id])
			if len(r.alive) == 0 {
				r.endGame()
			}
		}
	} else {
		// 给本人发送消息，减少护盾
		m := pb.BloodChangePush_30003{
			Index: index,
			Num:   attacked * (-1),
		}

		fmt.Println("attack----", m)
		r.playerStatus[index].player.SendMsg(&m, 30003)
	}
}

// 增加护甲
func (r *Room) addDef(player *Player, i int32) {
	r.mu.Lock()
	r.mu.Unlock()
	if !r.isStart(false) {
		return
	}
	player.addDef(i)
}

func (r *Room) addHp(player *Player, i int32) {
	r.mu.Lock()
	r.mu.Unlock()
	if !r.isStart(false) {
		return
	}

	added := player.addHp(i)
	if added > 0 {
		m := pb.BloodChangePush_30003{
			Index: int32(r.playerIndexMap[player.id]),
			Num:   added,
		}
		r.broadcast(&m, 30003)
	}

}

func (r *Room) allIndex() []int32 {
	a := make([]int32, 0, 0)
	for i := 0; i < r.posIndex; i++ {
		a = append(a, int32(r.pos[i]))
	}
	return a
}
func (r *Room) randIndex(except int) int {
	var i = rand.Intn(len(r.alive))
	for i == except {
		i = rand.Intn(len(r.alive))
	}

	return r.alive[i]
}

// 游戏结束
func (r *Room) endGame() {
	r.status = 2

}

func (r *Room) isEnd(lock bool) bool {
	if lock {
		r.mu.Lock()
		r.mu.Unlock()
	}

	return r.status == 2

}

func (r *Room) isInit(lock bool) bool{
	if lock {
		r.mu.Lock()
		r.mu.Unlock()
	}

	return r.status == 0
}

func createStatus(player *Player) *playerStatus {
	p := &playerStatus{player: player}
	return p
}
