package main

func main() {
	f(1)
}

func f(a int32) int32 {
	if a&1 == 1 {
		return 1
	} else {
		return 0
	}
}
