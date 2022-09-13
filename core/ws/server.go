package ws

import (
	"github.com/aisuosuo/letter/config"
	"github.com/aisuosuo/letter/config/log"
	"github.com/aisuosuo/letter/core/pb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	WsServer = &server{
		sessions:   make(map[string][]*session),
		message:    make(chan []byte),
		access:     make(chan *session),
		disconnect: make(chan *session),
	}
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type server struct {
	sessions   map[string][]*session //允许多设备同时连接
	message    chan []byte           //用户消息通知
	access     chan *session         //用户连接
	disconnect chan *session         //用户断开连接
}

func (t *server) Start() {
	for {
		select {
		case s := <-t.access:
			log.Logger.Infof("user %s access", s.name)
			t.sessions[s.name] = append(t.sessions[s.name], s)
		case s := <-t.disconnect:
			log.Logger.Infof("user %s disconnect", s.name)
			sessions := t.sessions[s.name]
			newSessions := make([]*session, 0)
			for _, item := range sessions {
				if item != s {
					newSessions = append(newSessions, s)
				}
			}
			t.sessions[s.name] = newSessions
		case data := <-t.message:
			message := new(pb.Message)
			err := proto.Unmarshal(data, message)
			if err != nil {
				log.Logger.Error(err)
				continue
			}
			if message.To == "" {
				log.Logger.Warnf("message not specify receiver :%#v", message)
				continue
			}
			if sessions := t.sessions[message.To]; len(sessions) > 0 {
				for _, s := range sessions {
					//发送给客户端
					s.writeCh <- data
				}
			}
		}
	}
}

func (t *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Error("upgrade:", err)
		return
	}
	//加载用户信息
	query := r.URL.Query()
	user, ok := query["user"]
	if !ok || len(user) == 0 {
		log.Logger.Error("invalid user")
		return
	}
	userName := user[0]

	s := &session{
		name:    userName,
		conn:    c,
		writeCh: make(chan []byte),
	}

	t.access <- s
	quit := make(chan bool)
	go s.read(quit)
	go s.write(quit)
	<-quit
}

func Run() {
	go WsServer.Start()
	wsAddr := config.GlobalConfig.Run.WebsocketAddr
	http.Handle("/letter/ws", WsServer)
	err := http.ListenAndServe(wsAddr, nil)
	if err != nil {
		log.Logger.Error(err)
		return
	}
}
