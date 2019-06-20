package main

import "fmt"

func square(n int) int {
	return n * n
}

func main() {

	sq := func(n int) int {
		return n * n
	}

	fmt.Printf("%T", sq)
	// fmt.Println(sq(2) == square(2))
}
