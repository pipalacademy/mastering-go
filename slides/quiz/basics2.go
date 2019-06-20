package main

import "fmt"

func foo(a int) int {
	a++
	return a
}

func main() {
	n := 2
	n = foo(n)
	fmt.Println(n)
}
