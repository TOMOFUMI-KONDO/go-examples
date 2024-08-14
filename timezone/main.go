package main

import (
	"fmt"
	"time"
)

func main() {
	// loc, _ := time.LoadLocation("Asia/Tokyo")
	loc := time.UTC
	t1 := time.Now().Add(9 * time.Hour)
	t2 := t1.In(loc)

	fmt.Println(t1)
	fmt.Println(t2)
}
