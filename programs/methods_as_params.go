package main

import (
	"fmt"
	"math"
)

/*

package main

import (
  "fmt"
)

type Foo int

func (f Foo) A() {
    fmt.Println("A")
}
func (f Foo) B() {
    fmt.Println("B")
}
func (f Foo) C() {
    fmt.Println("C")
}

func main() {
    var f Foo
    bar := func(foo func()) {
        f.foo()
    }
    bar(A)
    bar(B)
    bar(C)
}
*/

type myFloat float64

func (f myFloat) square() float64 {
	return float64(f * f)
}

func (f myFloat) sqrt() float64 {
	return math.Sqrt(float64(f))
}

func main() {
	floats := []float64{1.1, 2, 3, 4, 5}
	nums := make([]myFloat, len(floats))

	bar := func(fn func(f myFloat) float64) float64 {
		return f.fn()
	}

	for _, f := range nums {
		fmt.Println(bar(f.sqrt))
	}
}
