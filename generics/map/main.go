package main

import "fmt"

func main() {
	s := Map([]int{1, 2, 3, 4, 5}, func(v int) string { return fmt.Sprintf("%b", v) })
	fmt.Println(s)
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	mapped := make([]T2, 0, len(s))
	for _, v := range s {
		mapped = append(mapped, f(v))
	}
	return mapped
}
