package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/lucas-clemente/quic-go"
)

var (
	message = "hello"
	addr    string
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:44300", "server address")
	flag.Parse()
}

func main() {
	w, err := os.Create("keylog.txt")
	if err != nil {
		panic(err)
	}

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
		KeyLogWriter:       w,
	}

	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		panic(err)
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	fmt.Printf("send: '%s'\n", message)
	_, err = stream.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, len(message))
	_, err = io.ReadFull(stream, buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("got: '%s'\n", buf)
}
