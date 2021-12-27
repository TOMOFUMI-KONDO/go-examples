package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var host string

func init() {
	flag.StringVar(&host, "host", "localhost:8000", "server host")
	flag.Parse()
}

type Event struct {
	ID   string
	Data string
}

func main() {
	events, ctx, err := EventSource(fmt.Sprintf("http://%s/prime", host))
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-events:
			fmt.Printf("Event(Id=%s): %s\n", event.ID, event.Data)
		}
	}
}

func EventSource(url string) (chan Event, context.Context, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != 200 {
		return nil, nil, fmt.Errorf("Response Status Code should be 200, but %d\n", res.StatusCode)
	}

	ctx, cancel := context.WithCancel(req.Context())
	events := make(chan Event)
	go receiveSSE(events, cancel, res)

	return events, ctx, nil
}

func receiveSSE(events chan Event, cancel context.CancelFunc, res *http.Response) {
	reader := bufio.NewReader(res.Body)
	var buffer bytes.Buffer
	event := Event{}

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			close(events)
			if err == io.EOF {
				cancel()
				return
			}
			panic(err)
		}

		switch {
		case bytes.HasPrefix(line, []byte(":ok")):
			// skip
		case bytes.HasPrefix(line, []byte("id:")):
			event.ID = string(line[3 : len(line)-1])
		case bytes.HasPrefix(line, []byte("number:")):
			buffer.Write(line[7 : len(line)-1])
		case bytes.Equal(line, []byte("\n")):
			event.Data = buffer.String()
			buffer.Reset()
			if event.Data != "" {
				events <- event
			}
			event = Event{}
		default:
			fmt.Fprintf(os.Stderr, "Parse Error: %s\n", line)
			cancel()
			close(events)
		}
	}
}
