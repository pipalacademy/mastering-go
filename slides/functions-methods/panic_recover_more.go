package main

import "fmt"

func recovery() {
    if r := recover(); r != nil {
        fmt.Println("Recovered in f", r)
    }
}

func f() {
    defer recovery()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 2 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}

func main() {
    f()
    fmt.Println("Returned normally from f.")
}
