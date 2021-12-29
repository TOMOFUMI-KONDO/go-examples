package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"sync"

	"github.com/lucas-clemente/quic-go"
)

var (
	message = "hello"
	addr    string
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:4430", "server address")
	flag.Parse()
}

func main() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			session, err := quic.DialAddr(addr, tlsConf, nil)
			if err != nil {
				panic(err)
			}

			stream, err := session.OpenStreamSync(context.Background())
			if err != nil {
				panic(err)
			}
			defer stream.Close()

			fmt.Printf("%d send: '%s'\n", i, message)
			_, err = stream.Write([]byte(message))
			if err != nil {
				panic(err)
			}

			buf := make([]byte, len(message))
			_, err = io.ReadFull(stream, buf)
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}
			fmt.Printf("%d got: '%s'\n", i, buf)

			wg.Add(-1)
		}(i + 1)
	}

	wg.Wait()
}
