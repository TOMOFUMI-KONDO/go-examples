package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	timeOutContext, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(timeOutContext, "GET", "http://localhost:18888/slow-page", nil)
	if err != nil {
		panic(nil)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(nil)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
}
