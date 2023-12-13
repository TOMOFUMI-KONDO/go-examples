package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handleUpgrade(w http.ResponseWriter, r *http.Request) {
	myProtocolName := "MyProtocol"

	// this endpoint accepts only upgrade protocol
	if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != myProtocolName {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to " + myProtocolName)

	// get low layer socket
	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		panic(nil)
	}
	defer conn.Close()

	// send request for change protocol
	response := http.Response{StatusCode: 101, Header: make(http.Header)}
	response.Header.Set("Upgrade", myProtocolName)
	response.Header.Set("Connection", "Upgrade")
	response.Write(conn)

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(readWriter, "%d\n", i)
		fmt.Println("->", i)
		readWriter.Flush()
		receive, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("<- %s", string(receive))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/upgrade", handleUpgrade)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
