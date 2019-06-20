package main

import "fmt"

type A struct {
	n int
	f func()
}

func main() {
	a := A{}
	fmt.Println(a)
}
