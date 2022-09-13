package ws

import (
	"github.com/aisuosuo/letter/config/log"
	"github.com/gorilla/websocket"
)

type session struct {
	name    string
	conn    *websocket.Conn
	writeCh chan []byte
}

func (t *session) read(quit chan<- bool) {
	for {
		_, message, err := t.conn.ReadMessage()
		if err != nil {
			log.Logger.Error(err)
			t.conn.Close()
			close(t.writeCh)
			WsServer.disconnect <- t
			quit <- true
			return
		}
		WsServer.message <- message
	}
}

func (t *session) write(quit chan<- bool) {
	for {
		message := <-t.writeCh
		err := t.conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			log.Logger.Error(err)
			quit <- true
			return
		}
	}
}
