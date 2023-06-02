package ws

import (
	"github.com/aisuosuo/letter/config/log"
	"github.com/gorilla/websocket"
)

type session struct {
	uid     uint32
	conn    *websocket.Conn
	writeCh chan []byte
	quit    chan struct{}
}

func (t *session) read() {
	for {
		select {
		case <-t.quit:
			log.Logger.Debug("read recv quit")
			return
		default:
			_, message, err := t.conn.ReadMessage()
			if err != nil {
				log.Logger.Warn(err)
				WsServer.disconnect <- t
				return
			}
			WsServer.message <- message
		}

	}
}

func (t *session) write() {
	for {
		select {
		case <-t.quit:
			log.Logger.Debug("writer recv quit")
			return
		default:
			message := <-t.writeCh
			err := t.conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				log.Logger.Warn(err)
				return
			}
		}
	}
}
