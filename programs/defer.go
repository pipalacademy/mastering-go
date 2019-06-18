package main

import (
	"fmt"
)

func foo(s string) string {
	fmt.Println("1")
	return s
}

func hugeTask(s string) string {
	fmt.Println("2")
	return s
}

func main() {
	s := "go"
	defer foo(s)
	defer hugeTask(foo(s))
	fmt.Println("3")
}
