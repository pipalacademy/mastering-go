package main

import "fmt"

func changeFirst(a [5]int, b int) {
	a[0] = b
}
func main() {
	var a, b [5]int
	for i := range a {
		a[i] = i
	}
	fmt.Println(a, len(a))

	// pass by value or reference ?
	changeFirst(a, 42)
	fmt.Println(a)

	b = a
	fmt.Printf("%v %p %p %T\n", a == b, &a, &b, a)

	// does change in `a` reflect in `b`?
	a[2] = 100
	fmt.Println(a, b)
}
