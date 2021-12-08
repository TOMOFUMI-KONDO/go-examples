package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	l := Lock{}

	go l.lockWhile(3)

	fmt.Println(time.Now().Clock())
	l.exclusiveHello()
	fmt.Println(time.Now().Clock())
}

type Lock struct {
	m sync.Mutex
}

func (l *Lock) exclusiveHello() {
	l.lock()
	fmt.Println("hello")
	l.unlock()
}

func (l *Lock) lock() {
	l.m.Lock()
}

func (l *Lock) unlock() {
	l.m.Unlock()
}

func (l *Lock) lockWhile(n int) {
	l.lock()
	time.Sleep(time.Second * time.Duration(n))
	l.unlock()
}
