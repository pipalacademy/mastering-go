package main

import "fmt"

func main() {
	var a []int

	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a, len(a), cap(a))
}
