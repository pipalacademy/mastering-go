package main

import "fmt"

type Point struct{ x, y int }

func (p Point) setY(n int) {
	p.y = n
}

func (p *Point) setX(n int) {
	p.x = n
}

func main() {
	p := Point{1, 2}
	p.setX(50)
	p.setY(100)
	fmt.Println(p)
}
