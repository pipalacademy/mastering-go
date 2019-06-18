package main

import "fmt"

func main() {

	var m int // explicit type, only declaration
	m = 1

	var n int = 2                      // declaration along with initialization
	i := 3                             // short declarations
	var pi, e float32 = 3.14, 2.718281 // multiple variables with same type, multiple assignments
	var a, b = true, false

	// m = 5.6 		// compilation error
	fmt.Println(i, m, n, pi, e, a, b)

	var ( //grouping related variables
		name string = "alice"
		age  int    = 21
	)
	_, _ = name, age // ignore the values
}
