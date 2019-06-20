package main

import "fmt"

func f(s []int) {
	s = []int{4, 5, 6}
}

func g(s *[]int) {
	*s = []int{1, 2, 3}
}

func h(s []int) {
	s[0] = 100
}

func main() {
	var s []int
	f(s)
	fmt.Println(s, cap(s), len(s))
	g(&s)
	fmt.Println(s, cap(s), len(s))
	h(s)
	fmt.Println(s, cap(s), len(s))
}
