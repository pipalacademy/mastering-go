package main

import "fmt"

func main() {
	// this results in deadlock
	// As sender blocks until reciever recieves `1`.
	// As it's one single main go routine, the control never goes to next line,
	// so the reciever doesn't get chance to recieve it, making it a deadlock.
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
