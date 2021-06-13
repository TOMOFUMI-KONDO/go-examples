package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func jsonRequestHandler(w http.ResponseWriter, req *http.Request) {
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v\n", jsonBody)

	fmt.Fprintf(w, "<html><body><h1>Hello HTTP!</body></html>")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", jsonRequestHandler)
	log.Println("start http listening :80")
	httpServer.Addr = ":80"
	log.Println(httpServer.ListenAndServe())
}
