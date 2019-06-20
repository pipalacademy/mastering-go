package main

import (
	"fmt"
	"strings"
)

func mapper(f func(string) string, words []string) []string {
	r := make([]string, len(words))
	for i, w := range words {
		r[i] = f(w)
	}
	return r
}

func main() {
	s := []string{"one", "two", "three"}
	fmt.Println(mapper(strings.ToUpper, s)) //["ONE", "TWO", "THREE"]
}
