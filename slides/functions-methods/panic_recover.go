package main

import "fmt"

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func main() {
	defer recovery()
	var a []int
	fmt.Println(a[3])
	fmt.Println("Returning normally from main")
}
