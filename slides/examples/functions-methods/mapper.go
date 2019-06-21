package main

import (
	"fmt"
	"math"
)

func square(x float64) float64 {
	return x * x
}

func mapperFloat64(f func(float64) float64, alist []float64) []float64 {
	r := make([]float64, len(alist))
	for i, v := range alist {
		r[i] = f(v)
	}
	return r
}

func main() {
	nums := []float64{1.1, 2, 3, 4, 5}
	fmt.Println(mapperFloat64(square, nums))
	fmt.Println(mapperFloat64(math.Sqrt, nums))
}
