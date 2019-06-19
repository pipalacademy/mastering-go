package main

import "fmt"

func index(words []string, a string) int {
	for i, word := range words {
		if word == a {
			return i
		}
	}
	return -1
}

func main() {
	x := []string{"go", "python", "c"}
	fmt.Println(index(x, "go"))      // 0
	fmt.Println(index(x, "c"))       //2
	fmt.Println(index(x, "haskell")) // -1
}
