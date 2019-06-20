package main

import "fmt"

func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5), sum(100, 200))
	// Join("-", "foo", "bar", "chi") //foo-bar-chi
	// Join("-", "a", "b", "c", "d")  //a-b-c-d
}
