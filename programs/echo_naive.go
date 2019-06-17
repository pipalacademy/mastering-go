package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string = "", " "

	startIndex := 1
	if os.Args[1] == "-s" {
		sep = os.Args[2]
		startIndex = 3
	}
	for i := startIndex; i < len(os.Args)-1; i++ {
		s += os.Args[i] + sep
	}
	s += os.Args[len(os.Args)-1]
	fmt.Println(s)
}
