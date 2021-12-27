package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	director := func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = ":8000"
	}
	modifier := func(res *http.Response) error {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Reading body error: %w\n", err)
		}

		newBody := bytes.NewBuffer(body)
		newBody.WriteString("via Proxy\n")
		res.Body = ioutil.NopCloser(newBody)
		res.Header.Set("Content-Length", strconv.Itoa(newBody.Len()))

		return nil
	}

	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}
	server := http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: rp,
	}

	log.Println("Stat Listening at :8001")
	log.Println(server.ListenAndServe())
}
