package game

import (
	"fmt"
	"funygame/core"
	"funygame/pb"
	"funygame/utils"
	"sync"
)

// 游戏全局信息
type Game struct {
	ProcessMap map[int32]Process

	mu      sync.Mutex
	AddrMap map[string]*Player
	IdMap   map[int32]*Player
}

func (g *Game) EnterGame(msg EnterGameMsg, request *core.Request) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if p, ok := g.AddrMap[request.RemoteAddr]; ok {
		p.id = msg.P
		// 发送玩家进入的信息

	} else {
		core.Debug("没有用户的连接")
	}
}
func (g *Game) BroadcastMsg(msg interface{}, pid int32) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, v := range g.AddrMap {
		if v.id != pid {
			v.SendMsg(msg)
		}
	}
}

func (g *Game) RegisterPlayer(connection *core.Connection) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.AddrMap[connection.RemoteAddr()]; ok {
		core.Debug("%v 角色存在 ", connection.RemoteAddr())
	} else {
		g.AddrMap[connection.RemoteAddr()] = &Player{
			id:   0,
			conn: connection,
		}
	}
}

func (g *Game) CloseConnection(addr string) {
	if p, ok := g.AddrMap[addr]; ok {
		core.Debug("移除角色:%s,%v", addr, p.id)
		delete(g.AddrMap, addr)
		delete(g.IdMap, p.id)
	}
}

var GameVal = &Game{
	ProcessMap: make(map[int32]Process),
	AddrMap:    make(map[string]*Player),
	IdMap:      make(map[int32]*Player),
}

func Start() {

	InitProcess()
	sm := &core.ServeMux{}

	sm.HandleFunc("pf", func(r *core.Request, w *core.Response) {

		test := &pb.Message{}
		e := r.ReadPb(test)
		if e != nil {
			core.Debug(e.Error())
			return
		}
		fmt.Println(test)

		bytes := utils.MsgToBytes(test)
		w.Write(bytes)

	})

	s := &core.Server{
		Addr: "127.0.0.1:8900",
	}
	s.RegisterOnConnection(func(connection *core.Connection) {

		GameVal.RegisterPlayer(connection)
		core.Debug("地址：%s", connection.RemoteAddr())
	})

	s.RegisterOnClose(func(connection string) {
		core.Debug("关闭connec %v", connection)
		GameVal.CloseConnection(connection)
	})

	s.Handler = sm

	s.ListenAndServe()
}
