package main

import "fmt"

func main() {
	const pi = 3.14
	const secsPerDay = 60 * 60 * 24
	const avogadro = 6.02214129e23
	const planck = 6.62606957e-34

	fmt.Println(secsPerDay)
	// secsPerDay = 86400 // compilation error

	const ( // constants can also be grouped
		a = 1
		b = 1.23456
		c = true
	)
	fmt.Printf("%T %T %T %T %T", a, b, c, avogadro, planck)
}
