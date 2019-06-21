package main

import "fmt"

func main() {
	var x = []string{"go", "python", "c", "java", "haskell"} // slice

	y := x[0:3]
	z := x
	fmt.Printf("%T %v %v %v\n", y, y, cap(y), len(y))

	y[2] = "rust"
	fmt.Println(x, y, z)
}
