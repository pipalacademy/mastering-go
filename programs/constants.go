package main

import (
	"fmt"
)

func main() {
	const pi = 3.14
	const secsPerDay = 60 * 60 * 24

	fmt.Println(secsPerDay)
	// secsPerDay = 86400 // compilation error

	const ( // constants can be grouped
		a = 1
		b = 1.23456
		c = true
	)
	fmt.Println(a, b, c)
}
