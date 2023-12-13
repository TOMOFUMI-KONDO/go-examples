package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerWithCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>2回目以降</body></html>")
	} else {
		fmt.Fprintf(w, "<html><body>初めてのリクエスト</body></html>")
	}
}

func serve() {
	var httpServer http.Server
	http.HandleFunc("/", handlerWithCookie)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
