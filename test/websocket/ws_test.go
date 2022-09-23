package websocket

import (
	"fmt"
	"github.com/aisuosuo/letter/core/pb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	quit := make(chan bool)
	go ClientTest("wf", "ss")
	//go ClientTest("wf", "ss")
	//go ClientTest("ss", "wf")
	//go ClientTest("ss", "wf")
	<-quit
}

func ClientTest(user, to string) {
	u := fmt.Sprintf("ws://localhost:8086/letter/ws?user=%s", user)
	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		fmt.Println("dial:", err)
		return
	}

	defer c.Close()

	timer := time.NewTicker(5 * time.Second)

	//reader
	go func() {
		for {
			_, data, err := c.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			message := new(pb.Message)
			proto.UnmarshalMerge(data, message)
			fmt.Println(fmt.Sprintf("%s recv message:%v", user, message))
		}
	}()

	for {
		<-timer.C
		data := &pb.Message{
			From:    user,
			To:      to,
			Type:    1,
			Content: "hello",
		}
		message, _ := proto.Marshal(data)
		c.WriteMessage(websocket.BinaryMessage, message)
	}
}
