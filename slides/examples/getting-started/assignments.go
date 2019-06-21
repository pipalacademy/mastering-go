package main

import "fmt"

func main() {
	m, n := 1, 2
	m, n = m+1, n+1
	fmt.Println(m, n)

	m, n = n, m
	fmt.Println(m, n)
}
