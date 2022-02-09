package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"

	"github.com/lucas-clemente/quic-go"
)

var (
	addr string
)

func init() {
	addr = *flag.String("addr", ":44300", "server address")
	flag.Parse()
}

func main() {
	// make listener, specifying addr and tls config.
	// QUIC needs to be used with TLS.
	// see: https://www.rfc-editor.org/rfc/rfc9001.html
	listener, err := quic.ListenAddr(addr, genTLSConf(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("listening %s\n", addr)

	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			panic(err)
		}

		go func(sess quic.Session) {
			stream, err := sess.AcceptStream(context.Background())
			if err != nil {
				fmt.Printf("failed to accept stream: %s\n", err)
				return
			}
			defer stream.Close()

			// echo received data
			buf := make([]byte, 1024)
			nr, err := stream.Read(buf)
			if err == io.EOF {
				fmt.Println("connection closed.")
				return
			}
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(buf))
			if _, err := stream.Write(buf[0:nr]); err != nil {
				fmt.Println(err)
			}
		}(sess)
	}
}

func genTLSConf() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
