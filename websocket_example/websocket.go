package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"time"
)

func Echo(ws *websocket.Conn) {
	var err error

	count := 0
	for {
		<-time.After(time.Second * 1)
		count++
		if err = websocket.Message.Send(ws, count); err != nil {
			log.Println("Can't send")
			break
		}
	}
}
