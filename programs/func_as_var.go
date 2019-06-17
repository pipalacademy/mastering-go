package main

import "fmt"

func adder(m int) func(int) int {
	return func(n int) int {
		return m + n
	}
}

func main() {
	Add2 := adder(2)
	Add100 := adder(100)
	fmt.Println(Add2(1), Add100(1))
}
