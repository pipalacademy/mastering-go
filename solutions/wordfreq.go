// Write a program `wordfreq.go` to compute the frequency of all words in a file.
//
//     $ cat files/five.txt
//     five
//     five four
//     five four three
//     five four three two
//     five four three two one
//
//     $ go run wordfreq.go files/five.txt
//     2 two
//     1 one
//     5 five
//     4 four
//     3 three
//
// Please note that the order of lines in the output is not important.
//
// Hints:
//     - ioutil.ReadFile
//     - strings.Fields

package main

import (
	"fmt"
	"os"
	"io/ioutil"
    "log"
    "strings"
)

func main() {
    // take filename as command-line argument and read words from it
    filename := os.Args[1]
    contents, err := ioutil.ReadFile(filename)

    if err != nil {
        log.Fatal(err)
    }

    words := strings.Fields(string(contents))

    // map for storing the counts
    freq := make(map[string]int)

    // compute frequency
    for _, w := range words {
    	freq[w] += 1
    }

    // print frequency
    for w, count := range freq {
    	fmt.Printf("%d %s\n", count, w)
    }
}