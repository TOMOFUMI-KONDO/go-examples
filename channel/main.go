package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("start")

	c := make(chan int, 5)
	defer close(c)

	for i := 0; i < 5; i++ {
		go f(c)
	}

	for i := 0; i < 5; i++ {
		n := <-c
		fmt.Printf("got %d\n", n)
	}

}

func f(c chan int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1)
	c <- rand.Intn(10)
}
