package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a []int
	b := []int{}

	fmt.Println(a, b)
	fmt.Println(a == nil, b == nil)
	fmt.Println(cap(a), cap(b))
	fmt.Println(len(a), len(b))

	aj, _ := json.Marshal(a)
	bj, _ := json.Marshal(b)
	fmt.Printf("%s %s", aj, bj)
}
