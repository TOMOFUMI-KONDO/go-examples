package main

func main() {
	f(1)
}

func f(a int32) int32 {
	// if a&1 == 0
	if a%2 == 0 {
		return 1
	} else {
		return 0
	}
}
