package main

import (
    "fmt"
)

var count int

func display(name string, counter chan int, done chan bool) {
    for i := range counter {
        fmt.Printf("%d: %s\n", i, name)
    }
    done <- true
}

func makeCounter(counter chan int, n int) {
    for i := 0; i < n; i++ {
        counter <- i
    }
    close(counter)
}

func main() {
    done := make(chan bool)
    counter := make(chan int)

    go makeCounter(counter, 30)

    go display("A", counter, done)
    go display("B", counter, done)
    go display("C", counter, done)

    // block until done
    <-done
    <-done
    <-done
}
