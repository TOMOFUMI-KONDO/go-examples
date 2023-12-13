package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
)

func simpleRequest() {
	resp, err := http.Get("http://localhost:18888/chunked")
	if err != nil {
		panic(nil)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}
