package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var input string

func init() {
	flag.StringVar(&input, "i", "input.txt", "path of input file.")
	flag.Parse()
}

func main() {
	lines, err := load(input)
	if err != nil {
		log.Fatalf("failed to load input: %w", err)
	}

	// allocate enough memory space in advance.
	result := make([][32]byte, len(lines))

	var wg sync.WaitGroup
	wg.Add(len(lines))

	for i, line := range lines {
		go func(line []byte, i int) {
			defer wg.Done()
			result[i] = process(line)
		}(line, i)
	}

	// wait until all process for each line finishes.
	wg.Wait()

	output(result)
}

// load input file and return it as bytes-slice.
func load(path string) ([][]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	lines := make([][]byte, 0)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to ReadLine: %w", err)
		}

		lines = append(lines, line)
	}

	return lines, nil
}

// process a line of input.
func process(b []byte) [32]byte {
	return sha256.Sum256(b)
}

// output hex-dump of SHA256 checksum.
func output(bb [][32]byte) {
	for _, b := range bb {
		fmt.Printf("%x\n", b)
	}
}
