package main

import "iter"

func main() {
	for v:=Filter([]int{0,1,2,3,4,5}, func(v int)bool {return v%2==0}){
		fmt.Println(v)
	}
}

func Filter[T any](s []T, f func(v T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if !f(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}
