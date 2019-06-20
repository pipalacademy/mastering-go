package main

import "fmt"

func main() {
	var s []int
	s[0] = 1
	fmt.Println(s, cap(s), len(s))
}
