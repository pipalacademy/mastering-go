package main

import "fmt"

func main() {
	k := [2]int{1, 2}
	m := map[[2]int]string{k: "first"}
	fmt.Println(m) // map[[1 2]:first]
	k[1] = 100
	fmt.Println(m) // map[[1 2]:first]
}
