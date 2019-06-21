package main

import "fmt"

func foo(a int) {
	a++
}

func main() {
	n := 2
	foo(n)
	fmt.Println(n)
}
