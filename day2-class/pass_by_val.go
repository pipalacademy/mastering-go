package main

import "fmt"

func modify(n int) {
	n = n + 1
}

func main() {
	i := 2
	modify(i)
	fmt.Println(i)
}
