package game

import (
	"fmt"
	"funygame/core"
	"funygame/pb"
	"funygame/utils"
	"github.com/golang/protobuf/proto"
	"runtime"
	"sync"
)

var playerUid = core.AtomicInt64(0)

// 游戏全局信息
type Game struct {
	ProcessMap map[int32]Process

	mu      sync.Mutex
	AddrMap map[string]*Player
	IdMap   map[int64]*Player

	RoomManager *RoomManager
}

// 第一次连接进行玩家注册
func (g *Game) RegisterPlayer(connection *core.Connection) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.AddrMap[connection.RemoteAddr()]; ok {
		// 可能断线重连
		core.Debug("%v 角色存在 ", connection.RemoteAddr())
	} else {
		g.AddrMap[connection.RemoteAddr()] = &Player{
			id:      playerUid.AddAndGet(1),
			conn:    connection,
			msgChan: make(chan proto.Message),
			robot:   false,
			data:    &PlayerData{Hp:2000},
		}
	}
}

func (g *Game) CloseConnection(addr string) {
	if p, ok := g.AddrMap[addr]; ok {
		core.Debug("移除角色:%s,%v", addr, p.id)
		delete(g.AddrMap, addr)
		delete(g.IdMap, p.id)
		g.RoomManager.ExitRoom(p)
	}
}

var GameVal = &Game{
	ProcessMap:  make(map[int32]Process),
	AddrMap:     make(map[string]*Player),
	IdMap:       make(map[int64]*Player),
	RoomManager: CreateRoomManager(),
}

func Start() {

	InitProcess()
	sm := &core.ServeMux{}

	sm.HandleFunc("pb", func(r *core.Request, w *core.Response) {
		defer func() {
			if err := recover(); err != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				core.Debug("发生错误 %v####%s\n", err, buf)
			}
		}()
		msg := &pb.Message{}
		e := r.ReadPb(msg)

		if e != nil {
			core.Debug(e.Error())
			return
		}

		p, _ := GameVal.AddrMap[r.RemoteAddr]

		if p == nil {
			core.Error("用户不存在:%v", r.RemoteAddr)
			return
		}

		if v, ok := GameVal.ProcessMap[msg.MsgNo]; ok {
			var msgBody proto.Message
			if msg.Body != nil {
				message := v.Msg()
				e := proto.Unmarshal(msg.Body, message)
				if e != nil {
					core.Debug("解析body出错: %s", e.Error())
					return
				}
				core.Logf("收到消息:%v", message)
				msgBody = message;
			}

			retMsg := v.Action(msgBody, p)
			if retMsg != nil {
				b, _ := proto.Marshal(retMsg)
				m := &pb.Message{
					Seq:   msg.Seq,
					MsgNo: msg.MsgNo,
					Body:  b,
					Uid:   0,
				}
				fmt.Printf("[%v]响应消息，消息号:%v\n",p.id,msg.MsgNo)
				bytes := utils.MsgToBytes(m)
				w.Write(bytes)
			}
		} else {
			core.Error("消息不存在:%v", msg.MsgNo)
			return
		}

	})

	s := &core.Server{
		Addr: "127.0.0.1:8900",
	}
	s.RegisterOnConnection(func(connection *core.Connection) {

		GameVal.RegisterPlayer(connection)
		core.Debug("注册地址：%s", connection.RemoteAddr())
	})

	s.RegisterOnClose(func(connection string) {
		core.Debug("关闭connec %v", connection)
		GameVal.CloseConnection(connection)
	})

	s.Handler = sm

	s.ListenAndServe()
}
