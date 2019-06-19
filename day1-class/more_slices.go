package main

import "fmt"

func changeFirst(a []int, b int) {
	a[0] = b
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	changeFirst(a, 42)
	fmt.Println(a)
}
