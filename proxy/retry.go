package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	target, err := url.Parse("http://127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	rp := httputil.NewSingleHostReverseProxy(target)
	rp.Transport = &RetryTransport{}
	server := http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: rp,
	}

	log.Println("Start Listening at :8001")
	log.Println(server.ListenAndServe())
}

type RetryTransport struct {
}

func (RetryTransport) RoundTrip(r *http.Request) (resp *http.Response, err error) {
	for i := 0; i < 3; i++ {
		resp, err := http.DefaultTransport.RoundTrip(r)
		if err != nil {
			log.Printf("failed to RoundTrip: %v\n", err)
			time.Sleep(time.Second)
			continue
		}
		return resp, nil
	}

	return nil, fmt.Errorf("failed to request to %s", r.URL.String())
}
