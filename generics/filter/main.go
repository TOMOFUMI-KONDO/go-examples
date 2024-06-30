package main

import "fmt"

func main() {
	s := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
	fmt.Println(s)
}

func Filter[T any](s []T, f func(v T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range s {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
