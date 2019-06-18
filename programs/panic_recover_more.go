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

/*
Calling g.
printing in g, 0
printing in g, 1
printing in g, 2
Panicking!
defer in g, 2
defer in g, 1
defer in g, 0
Recovered in f, 3
Returned normally from f.
*/
