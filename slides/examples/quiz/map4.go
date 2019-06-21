package main

import "fmt"

func foo(m map[int]int) {
	for _, v := range m {
		v++
	}
}

func main() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println(m)
}
