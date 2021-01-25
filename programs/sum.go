package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	for i := 1; i < len(os.Args); i++ {
		val, err := strconv.Atoi(os.Args[i])
		if err != nil {
			panic(fmt.Sprintf("Invalid number %s", os.Args[i]))
		}
		sum += val
	}
	fmt.Println(sum)
}
