package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	port  = 18443
	image []byte
)

func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

func serve() {
	http.HandleFunc("/", handleHtml)
	http.HandleFunc("/image", handleImage)

	fmt.Printf("start http listening :%d\n", port)
	fmt.Println(http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "../tls/server.crt", "../tls/server.key", nil))
}

func handleHtml(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/image", nil)
	}

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image" /></body></html>`)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}
