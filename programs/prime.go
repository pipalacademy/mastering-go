package main

import "fmt"

func isPrime(n int) bool {
	if n < 0 {
		n = -n
	}
	if n < 2 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 0; i < 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
