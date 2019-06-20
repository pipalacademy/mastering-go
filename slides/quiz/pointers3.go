package main

import "fmt"

func foo(p *int) {
	m := 1
	p = &m
}

func main() {
	n := 2
	p := &n
	*p = 5
	foo(p)
	fmt.Println(n)
}
