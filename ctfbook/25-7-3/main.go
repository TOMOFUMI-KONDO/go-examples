package main

func main() {
	f(10)
}

func f(a int32) int32 {
	var b int32 = 0
	var c int32 = 1

	for c <= a {
		b += c
		c += 1
	}
	return b
}
