package ws

import (
	"github.com/aisuosuo/letter/api/service"
	"github.com/aisuosuo/letter/config"
	"github.com/aisuosuo/letter/config/log"
	"github.com/aisuosuo/letter/core/pb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	WsServer = &server{
		sessions:   make(map[uint32][]*session),
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
	sessions   map[uint32][]*session //允许多设备同时连接
	message    chan []byte           //用户消息通知
	access     chan *session         //用户连接
	disconnect chan *session         //用户断开连接
}

func (t *server) Start() {
	timer := time.NewTicker(60 * time.Second)
	for {
		select {
		case s := <-t.access:
			log.Logger.Infof("user %d access", s.uid)
			t.sessions[s.uid] = append(t.sessions[s.uid], s)
		case s := <-t.disconnect:
			log.Logger.Infof("user %d disconnect", s.uid)
			sessions := t.sessions[s.uid]
			newSessions := make([]*session, 0)
			for _, item := range sessions {
				if item != s {
					newSessions = append(newSessions, item)
				}
			}
			//已经被移除了
			if len(sessions) == len(newSessions) {
				continue
			}

			if len(newSessions) == 0 {
				delete(t.sessions, s.uid)
			} else {
				t.sessions[s.uid] = newSessions
			}
			close(s.quit)
			close(s.writeCh)
			s.conn.Close()
			s = nil //回收
		case data := <-t.message:
			message := new(pb.Message)
			err := proto.Unmarshal(data, message)
			if err != nil {
				log.Logger.Error(err)
				continue
			}
			log.Logger.Debugf("recv:%#v", message)
			if message.To == 0 {
				log.Logger.Warnf("message not specify receiver :%#v", message)
				continue
			}
			err = service.UserService.AddMessage(uint(message.From), uint(message.To), message.Content)
			if err != nil {
				log.Logger.Warnf("insert to db err:%s", err.Error())
			}
			if sessions := t.sessions[message.To]; len(sessions) > 0 {
				for _, s := range sessions {
					go func(s *session) {
						defer func() {
							if err := recover(); err != nil {
								log.Logger.Error(err)
								t.disconnect <- s
							}
						}()
						//发送给客户端
						s.writeCh <- data
					}(s)
				}
			}
		case <-timer.C:
			log.Logger.Debugf("current session %v", t.sessions)
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
	user := strings.TrimPrefix(r.URL.Path, "/letter/ws/")
	uid, err := strconv.Atoi(user)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	s := &session{
		uid:     uint32(uid),
		conn:    c,
		writeCh: make(chan []byte),
		quit:    make(chan struct{}),
	}

	t.access <- s
	go s.read()
	go s.write()
	<-s.quit
}

func Run() {
	go WsServer.Start()
	wsAddr := config.GlobalConfig.RunConfig.WebsocketAddr
	http.Handle("/letter/ws/", WsServer)
	err := http.ListenAndServe(wsAddr, nil)
	if err != nil {
		log.Logger.Error(err)
		return
	}
}
