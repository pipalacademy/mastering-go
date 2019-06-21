package main

import "fmt"

func changeFirst(a [5]int, b int) {
	a[0] = b
}
func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// pass by value or reference ?
	changeFirst(a, 42)
	fmt.Println(a)
}
