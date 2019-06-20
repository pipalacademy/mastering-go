package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{0, 0}

	fmt.Println(Distance(p, q))
	fmt.Println(q)
	fmt.Println(p.DistanceFrom(q))

}
