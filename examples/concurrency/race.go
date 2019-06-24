package main

import (
    "fmt"
    "sync"
)

var sum int
var lock sync.Mutex

func compute(n int, w *sync.WaitGroup) {
    for i := 0; i < n; i++ {
        //lock.Lock()
        sum++
        //lock.Unlock()
    }
    w.Done()
}

func main() {
    var w sync.WaitGroup

    w.Add(3)
    n := 100000
    go compute(n, &w)
    go compute(n, &w)
    go compute(n, &w)
    w.Wait()
    fmt.Println("sum =", sum)
}
