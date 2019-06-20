package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

type Line struct {
	begin Point
	end   Point
}

type Circle struct {
	Point
	Line
	Radius float32
}

// func (l Line) Distance(q Point) float64 {
// 	return 100.0
// }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// func (c Circle) Distance(q Point) float64 {
// 	return c.Point.Distance(q) - float64(c.Radius)
// }

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	l := Line{Point{0, 0}, Point{1, 1}}
	c := Circle{Point{1, 2}, l, 2}
	// fmt.Println(c)
	// fmt.Println(c.X)
	fmt.Println(c.Line.Distance(Point{0, 0}))
	// fmt.Println(c.Point.Distance(Point{0, 0}))
}
