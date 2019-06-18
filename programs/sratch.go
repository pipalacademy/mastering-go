package main

import (
	"fmt"
)

func c() (i int) {
	defer func() { i++ }()
	return 100
}

func main() {
	r := `a multi
        line
        string`

	// var f float64
	// var i uint

	// for i = 0; i < 32; i++ {
	// 	f = float64(uint(1) << i)
	// 	if f == f+1 {
	// 		fmt.Println(i)
	// 	}
	// }
	// fmt.Println(strings.ToUpper("one"))
	// fmt.Println(strings.Repeat("a", 5))
	fmt.Println(c(), r)
}
