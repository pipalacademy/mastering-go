package main

import "fmt"

func main() {
	a := make([]int, 10)

	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a, cap(a), len(a))
}
