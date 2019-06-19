package main

import (
	"fmt"
	"os"
	"strconv"
)

// hint: use strconv.Atoi

func main() {
	arg := os.Args[1]
	n, _ := strconv.Atoi(arg)
	fmt.Println(n * n)
}
