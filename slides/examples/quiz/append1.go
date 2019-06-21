package main

import "fmt"

func main() {
	s := make([]string, 5)
	s = append(s, "a")
	fmt.Printf("%v %d %d\n", s, len(s), cap(s))
}
