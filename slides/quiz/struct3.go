package main

import "fmt"

func (n int) prInt() {
	fmt.Println(n)
}

func main() {
	m := 2
	m.prInt()
}
