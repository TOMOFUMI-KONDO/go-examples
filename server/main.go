package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8000", "listen port")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", handle)

	fmt.Printf("listen http %s\n", port)
	http.ListenAndServe("localhost"+port, nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
