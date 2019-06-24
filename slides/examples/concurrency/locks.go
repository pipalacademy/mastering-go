package main

import (
    "fmt"
    "sync"
)

var count int
var lock sync.Mutex

func compute(n int, done chan bool) {
    for i := 0; i < n; i++ {
        // protect the critical section with lock
        lock.Lock()
        count++
        lock.Unlock()
    }
    done <- true
}

func main() {
    done := make(chan bool)
    n := 1000

    go compute(n, done)
    go compute(n, done)
    go compute(n, done)

    // block until done
    <-done
    <-done
    <-done

    fmt.Println(count)
}
