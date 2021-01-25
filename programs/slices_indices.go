package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a[1:2])
	fmt.Println(a[:8])
	fmt.Println(a[1:])
	fmt.Println(a[1:6])
}
