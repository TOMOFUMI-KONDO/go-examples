package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("テキスト")

	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
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
