package main

import "fmt"

func main() {
	n := 2
	p := &n
	*p = 5
	fmt.Println(n)
}
