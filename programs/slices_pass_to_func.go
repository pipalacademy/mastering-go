package main

import "fmt"

func changeFirst(a []int, b int) {
	a[0] = b
}

func main() {
	a := []int{1, 2, 3, 4}

	fmt.Println(a)
	changeFirst(a, 100)
	fmt.Println(a)
}
