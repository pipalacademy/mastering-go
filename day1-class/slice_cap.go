package main

import "fmt"

func main() {
	var c []int // nil
	// c[0] = 1		// would err out with index out of range
	c = append(c, 1)
	fmt.Printf("%#v", c)
}
