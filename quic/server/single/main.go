package main

import (
	"context"
	"flag"
	"fmt"
	tls "github.com/TOMOFUMI-KONDO/go-sandbox/quic/server"
	"io"

	"github.com/lucas-clemente/quic-go"
	//"github.com/TOMOFUMI-KONDO/go-sandbox/quic/server/tls"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":4430", "server address")
	flag.Parse()
}

func main() {
	// make listener, specifying addr and tls config.
	// QUIC needs to be used with TLS.
	// see: https://www.rfc-editor.org/rfc/rfc9001.html
	listener, err := quic.ListenAddr(addr, tls.GenerateTLSConfig(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("listening %s\n", addr)

	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			panic(err)
		}

		go func() {
			stream, err := sess.AcceptStream(context.Background())
			if err != nil {
				fmt.Printf("failed to accept stream: %s\n", err)
				return
			}
			defer stream.Close()

			// echo received data
			buf := make([]byte, 1024)
			nr, err := stream.Read(buf)
			if err != nil && err == io.EOF {
				fmt.Printf("failed to read stream: %s\n", err)
				return
			}

			if nr < 1 {
				return
			}
			if _, err := stream.Write(buf[0:nr]); err != nil {
				fmt.Printf("failed to write stream: %s\n", err)
				return
			}
			fmt.Println(string(buf))
		}()
	}
}
