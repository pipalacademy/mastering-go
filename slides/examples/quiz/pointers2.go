package main

import "fmt"

func foo(p *int) {
	*p = 100
}

func main() {
	n := 2
	foo(&n)
	fmt.Println(n)
}
