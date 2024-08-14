package main

import "fmt"

func main() {
	for v := range Seq([]int{1, 3, 5}) {
		fmt.Println(v)
	}
}

func Seq[T any](s []T) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}
