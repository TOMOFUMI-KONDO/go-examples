package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8000", "listen port")
}

func main() {
	http.HandleFunc("/", echo)

	fmt.Printf("websocket listen localhost%s\n", port)
	fmt.Println(http.ListenAndServe("localhost"+port, nil))
}

var upgrader = websocket.Upgrader{} //use default options

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to Upgrade: %v\n", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Printf("failed to ReadMessage: %v\n", err)
			return
		}
		if messageType != websocket.TextMessage {
			log.Printf("messageType: %d is not supported\n", messageType)
			return
		}

		fmt.Printf("received: %s\n", string(p))

		if err := conn.WriteMessage(websocket.TextMessage, p); err != nil {
			log.Printf("failed to WriteMessage: %v\n", err)
			return
		}
	}
}
