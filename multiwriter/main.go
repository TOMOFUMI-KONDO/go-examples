package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

var src = map[string]string{
	"Hello": "World",
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	gzipWriter := gzip.NewWriter(w)
	gzipWriter.Header.Name = "sample.json"
	defer gzipWriter.Close()

	multiWriter := io.MultiWriter(gzipWriter, os.Stdout)

	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(src); err != nil {
		panic(err)
	}

	if err := gzipWriter.Flush(); err != nil {
		panic(err)
	}
}
