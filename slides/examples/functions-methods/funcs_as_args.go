package main

import (
	"fmt"
	"time"
)

func runAfter(g func(string), s string, milliSecs time.Duration) {
	time.Sleep(milliSecs * time.Millisecond)
	g(s)
}

func print(s string) {
	fmt.Println(s)
}

func main() {
	s := "hello"
	runAfter(print, s, 1000)
}
