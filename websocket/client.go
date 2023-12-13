package main

import (
	"flag"
	"fmt"

	"github.com/gorilla/websocket"
)

var host string

func init() {
	flag.StringVar(&host, "host", "ws://localhost:8000", "server host")
	flag.Parse()
}

func req() {
	conn, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		panic(err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		panic(err)
	}

	messageType, p, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}

	fmt.Printf("messageType: %d, data: %s\n", messageType, string(p))
}
