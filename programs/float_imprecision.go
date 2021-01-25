package main

import "fmt"

func main() {
	var f float32 = 1 << 24
	var i uint = 1 << 24

	fmt.Println(f == f+1)
	fmt.Println(i == i+1)

	for i = 0; i < 32; i++ {
		f = float32(uint(1) << i) // 1 << i
		if f == f+1 {
			fmt.Println(i)
		}
	}
}
