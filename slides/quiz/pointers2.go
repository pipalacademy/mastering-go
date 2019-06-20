package main

import "fmt"

func foo(p *int) {
	*p = 100
}

func main() {
	n := 2
	p := &n
	*p = 5
	foo(p)
	fmt.Println(n)
}
