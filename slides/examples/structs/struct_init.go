package main

import "fmt"

type Point struct {
	x int
	y int
}

// type Point struct{ x, y int }

func main() {
	p := Point{1, 2}

	var q Point
	q.x = 1

	r := Point{y: 1}

	fmt.Println(p, q, r)

}
