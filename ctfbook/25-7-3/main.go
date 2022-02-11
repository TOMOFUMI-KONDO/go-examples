package main

func main() {
	f(10)
}

func f(a int32) int32 {
	var b int32 = 0

	for i := 1; int32(i) <= a; i++ {
		b += int32(i)
	}
	return b
}
