package main

import (
	"fmt"
	"math"
)

type Point struct{ x, y float64 }

type Circle struct {
	Point
	radius int
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (c Circle) Distance(q Point) float64 {
	return c.Point.Distance(q) - float64(c.radius)
}

func main() {
	p := Point{2, 2}
	q := Point{0, 0}
	c := Circle{p, 1}

	fmt.Println(p.Distance(q))
	fmt.Println(c.Distance(q))
}
