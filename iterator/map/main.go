package main

import (
	"fmt"
	"iter"
)

func main() {
	for v := range Map([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(v int) string { return fmt.Sprintf("%b", v) }) {
		fmt.Println(v)
	}
}

func Map[T1, T2 any](s []T1, f func(v T1) T2) iter.Seq[T2] {
	return func(yield func(T2) bool) {
		for _, v := range s {
			if !yield(f(v)) {
				return
			}
		}
	}
}
