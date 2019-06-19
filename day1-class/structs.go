package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)
}
