package main

import "fmt"

func main() {
	var p *int

	a := 1
	p = &a

	fmt.Println(p)
}
