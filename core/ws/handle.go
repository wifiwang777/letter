package ws

import (
	"github.com/aisuosuo/letter/config/log"
	"github.com/gorilla/websocket"
)

type session struct {
	uid     uint32
	conn    *websocket.Conn
	writeCh chan []byte
}

func (t *session) read(quit chan bool) {
	for {
		select {
		case <-quit:
			log.Logger.Debug("writer recv quit")
			return
		default:
			_, message, err := t.conn.ReadMessage()
			if err != nil {
				log.Logger.Error(err)
				WsServer.disconnect <- t
				quit <- true
				return
			}
			WsServer.message <- message
		}

	}
}

func (t *session) write(quit chan bool) {
	for {
		select {
		case <-quit:
			log.Logger.Debug("writer recv quit")
			return
		default:
			message := <-t.writeCh
			err := t.conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				log.Logger.Error(err)
				WsServer.disconnect <- t
				quit <- true
				return
			}
		}
	}
}
