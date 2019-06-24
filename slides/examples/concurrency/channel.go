package main

import "fmt"

func send(ch chan int, num int) {
	ch <- num
}

func main() {
	ch := make(chan int)
	go send(ch, 1)
	fmt.Println(<-ch)
}
