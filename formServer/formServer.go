package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Println(string(dump))

	err = r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Println("gachas:" + r.Form.Get("gachas"))
	fmt.Println("leastTimes:" + r.Form.Get("leastTimes"))
	fmt.Println("maxCost:" + r.Form.Get("maxCost"))
	fmt.Println("mode:" + r.Form.Get("mode"))

	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "200")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", postFormHandler)
	log.Println("start http listening :80")
	httpServer.Addr = ":80"
	log.Println(httpServer.ListenAndServe())
}
