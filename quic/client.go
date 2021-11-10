package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"

	"github.com/lucas-clemente/quic-go"
)

func main() {
	if err := clientMain(); err != nil {
		log.Fatal(err)
	}
}

func clientMain() error {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		return err
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}

	message := "hello"
	fmt.Printf("Client: Sending '%s'\n", message)
	_, err = stream.Write([]byte(message))
	if err != nil {
		return err
	}

	buf := make([]byte, len(message))
	_, err = io.ReadFull(stream, buf)
	if err != nil {
		return err
	}
	fmt.Printf("Client: Got '%s'\n", buf)

	return nil
}
