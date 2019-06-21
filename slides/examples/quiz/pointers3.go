package main

import "fmt"

func foo(p *int) {
	m := 1
	p = &m
}

func main() {
	n := 2
	p := &n
	foo(p)
	fmt.Println(n)
}
