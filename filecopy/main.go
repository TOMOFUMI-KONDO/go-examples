package main

import (
	"flag"
	"io"
	"os"
)

var (
	old, new string
)

func init() {
	flag.StringVar(&old, "old", "old.txt", "old file name")
	flag.StringVar(&new, "new", "new.txt", "new file name")
	flag.Parse()
}

func main() {
	oldFile, err := os.Open(old)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create(new)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, oldFile); err != nil {
		panic(err)
	}
}
