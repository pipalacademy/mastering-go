package main

import "fmt"

func foo(m map[int]int) {
	for k := range m {
		m[k] = 42
	}
}

func main() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	foo(m)
	fmt.Println(m)
}
