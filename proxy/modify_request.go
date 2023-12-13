package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func ServeModifyReq() {
	director := func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = ":8000"
	}
	rp := &httputil.ReverseProxy{
		Director: director,
	}
	server := http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: rp,
	}
	log.Println("Start listening at :8001")
	log.Fatalln(server.ListenAndServe())
}
