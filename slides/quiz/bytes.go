package main

import "fmt"

func main() {
	s := []byte("hello")
	s[0] = byte('H')
	fmt.Println(string(s))
}
