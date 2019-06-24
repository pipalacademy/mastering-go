package main

import "fmt"

func numbers(n int) <-chan int {
    c := make(chan int)
    go func() {
        for i := 0; i < n; i++ {
            c <- i
        }
        close(c)
    }()
    return c
}

func squares(seq <-chan int) <-chan int {
    c := make(chan int)
    go func() {
        for n := range seq {
            c <- n * n
        }
        close(c)
    }()
    return c
}

func printer(seq <-chan int) {
    for n := range seq {
        fmt.Println(n)
    }
}

func main() {
    n := numbers(10)
    sq := squares(n)
    printer(sq)
}
