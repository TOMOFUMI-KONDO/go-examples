package main

import (
	"encoding/csv"
	"flag"
	"os"
	"time"
)

var delimiter string

func init() {
	flag.StringVar(&delimiter, "delimiter", "", "csv delimiter")
	flag.Parse()
}

func main() {
	f, err := os.Create("csv.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if delimiter != "" {
		w.Comma = []rune(delimiter)[0]
	}
	w.Write([]string{"id", "title", "created_at"})
	w.Write([]string{"1", "super item", time.Now().String()})
	w.Flush()
}
