package main

import (
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()
}

func main() {
	http.HandleFunc("/prime", handlePrimeSSE)

	fmt.Printf("start http listening :%d", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handlePrimeSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
	}

	// get context to detect disconnection
	ctx := r.Context()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var num int64
	for id := 1; id <= 10; id++ {
		// end when disconnect
		select {
		case <-ctx.Done():
			fmt.Println("Connection closed from client")
			return
		default:
			// do nothing
		}

		for {
			num++
			if big.NewInt(num).ProbablyPrime(20) {
				fmt.Printf("id: %d, number:%d\n", id, num)
				fmt.Fprintf(w, "id: %d\n", id)
				fmt.Fprintf(w, "number: %d\n\n", num)
				flusher.Flush()
				time.Sleep(time.Second)
				break
			}
		}
		time.Sleep(time.Second)
	}

	fmt.Println("Connection closed from server")
}
