package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	client := &http.Client{}
	readFile, err := os.Open("text.txt")
	if err != nil {
		panic(nil)
	}

	request, err := http.NewRequest("POST", "http://localhost:18888", readFile)
	if err != nil {
		panic(nil)
	}

	request.Header.Add("Content-Type", "text/plain")

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
