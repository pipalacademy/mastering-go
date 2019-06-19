// Write a program dups.go to find and print all the duplicate words in a sentence.
//
//     $ go run dups.go one two three two one
//     two
//     one
//
//     $ go run dups.go one two one two one
//     one
//     two
package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    seen := make(map[string]int)

    for _, a := range args {
        // print it when the duplicate is seen identified for the first time
        if seen[a] == 1 {
            fmt.Println(a)
        }
        seen[a]++
    }
}