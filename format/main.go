package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("format.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "string: %s\n", "string")
	fmt.Fprintf(f, "int: %d\n", 1)
	fmt.Fprintf(f, "float: %f\n", 1.1)
}
