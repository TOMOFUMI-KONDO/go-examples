package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("get/get.go")
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
}
