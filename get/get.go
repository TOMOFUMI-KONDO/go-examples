package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(nil)
	}

	resp, err := client.Do(request)
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
