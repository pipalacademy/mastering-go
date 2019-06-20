package main

import "fmt"

func (n int) foo() {
	fmt.Println(n)
}

func main() {
	n := 2
	n.foo()
}
