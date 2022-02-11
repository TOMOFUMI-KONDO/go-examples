package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening :8000")

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	var total int64
	buf := make([]byte, 1024)
	for {
		nr, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		total += int64(nr)
		fmt.Printf("received %dbyte\nnow total %dbyte\n", nr, total)
	}
}
