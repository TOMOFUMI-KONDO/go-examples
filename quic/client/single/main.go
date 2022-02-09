package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lucas-clemente/quic-go"
)

var (
	addr string
)

func init() {
	addr = *flag.String("addr", "localhost:44300", "server address")
	flag.Parse()
}

func main() {
	w, err := os.Create("keylog.txt")
	if err != nil {
		panic(err)
	}

	session, err := quic.DialAddr(addr, genTLSConf(w), nil)
	if err != nil {
		log.Fatalln(err)
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.Close()

	fmt.Println("send: hello")
	_, err = stream.Write([]byte(("hello")))
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, len("hello"))
	_, err = io.ReadFull(stream, buf)
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("receive: %s\n", buf)
}

func genTLSConf(w io.Writer) *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
		KeyLogWriter:       w,
	}
}
