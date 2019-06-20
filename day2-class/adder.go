package main

import "fmt"

func adder(n int) func(int) int {
	return func(m int) int {
		return m + n
	}
}

func main() {
	add2 := adder(2)
	add100 := adder(100)
	fmt.Println(add2(1), add100(1)) //	3, 101
}
