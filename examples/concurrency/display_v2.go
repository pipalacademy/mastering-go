package main

import (
    "fmt"
)

func display(name string, n int, done chan bool) {
    for i := 0; i < n; i++ {
        fmt.Printf("%d: %s\n", i, name)
    }
    done <- true
}

func main() {
    done := make(chan bool)

    go display("A", 10, done)
    go display("B", 10, done)
    go display("C", 10, done)

    // receive something from the channel. blocks until someone sends it
    <-done
    <-done
    <-done
}
