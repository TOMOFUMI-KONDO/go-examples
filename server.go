package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func dumpRequest(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Println(string(dump))
}

func handler(w http.ResponseWriter, r *http.Request) {
	dumpRequest(w, r)

	fmt.Fprintf(w, "<html><body><h1>Hello HTTP!</body></html>\n")
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	dumpRequest(w, r)

	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body><h1>2回目以降のリクエスト</body></html>\n")
	} else {
		fmt.Fprintf(w, "<html><body><h1>初めてのリクエスト</body></html>\n")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/cookie", cookieHandler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
