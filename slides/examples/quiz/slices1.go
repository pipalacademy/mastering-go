package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Println(s, cap(s), len(s))
	s = append(s, 1)
	fmt.Println(s, cap(s), len(s))
}
