package main

import (
	"code.google.com/p/go.net/websocket"
	"io/ioutil"
	"log"
	"time"
)

func Echo(ws *websocket.Conn) {
	filename := "echo.txt"
	for {
		<-time.After(time.Second * 1)

		buf, err := ioutil.ReadFile(filename)
		if err != nil {
			continue
		}

		if err = websocket.Message.Send(ws, string(buf)); err != nil {
			log.Println("Can't send")
			break
		}
	}
}
