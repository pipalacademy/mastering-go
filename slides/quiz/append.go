package main

import "fmt"

func main() {
	var s []string
	s = append(s, "a")
	fmt.Println(s, len(s), cap(s))
}
