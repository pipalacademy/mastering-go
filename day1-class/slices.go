package main

import "fmt"

func main() {
	s := []string{"go", "workshop"}

	s = []string{"go", "workshop", "day1"}
	fmt.Println(s)

	c := byte(65)

	fmt.Printf("%T %v %q", c, c, c)

}
