package main

import "fmt"

func f(s []int) {
	s = []int{1, 2, 3}
}

func g(s *[]int) {
	*s = []int{1, 2, 3}
}

func main() {
	var s []int
	f(s)
	fmt.Println(s, cap(s), len(s))
	g(&s)
	fmt.Println(s, cap(s), len(s))
}
