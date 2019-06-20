package main

import "fmt"

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func f() {
	defer recovery()
	var s []int
	s[0] = 1
}

func main() {
	f()
	fmt.Println("fasf")
}
