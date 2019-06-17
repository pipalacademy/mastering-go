package main

import (
	"fmt"
	"math"
)

func main() {

	var m int // explicit type, only declaration
	m = 1

	var n int = 2 // declaration along with initialization

	i := 3 // short declarations

	var pi, e float32 = 3.14, 2.718281 // multiple variables with same type, multiple assignments

	var a, b = true, false

	// m = 5.6 		// compilation error
	fmt.Println(i, m, n, pi, e, a, b)

	var (
		firstname string
		lastname  string
		age       int
	)
	fmt.Println(firstname, lastname, age)

	var (
		x float64
		y float64
		z float64
	)
	z = math.Sqrt(x*x + y*y)
	fmt.Println(x, y, z)
}
