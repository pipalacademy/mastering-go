package main

import (
	"fmt"
	"time"
)

func display(name string, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d: %s\n", i, name)
	}
}

func main() {
	go display("A", 10)
	display("B", 10)

	// hack to wait till the goroutine finished
	time.Sleep(1 * time.Second)
}
