package main

import (
	"flag"
	"io"
	"io/ioutil"
	"net"
	"os"
)

var size int64

func init() {
	flag.Int64Var(&size, "size", 1e9, "data size to be sent")
	flag.Parse()
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	file, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	if err := file.Truncate(size); err != nil {
		panic(err)
	}

	var offset int64 = 0
	buf := make([]byte, 1024)
	for {
		nr, errRead := file.ReadAt(buf, offset)
		if errRead != nil && errRead != io.EOF {
			panic(err)
		}

		offset += int64(nr)
		if _, err := conn.Write(buf[:nr]); err != nil {
			panic(err)
		}

		if errRead == io.EOF {
			break
		}
	}
}
