package main

import "fmt"

func main() {
	m := map[string]int{"zero": 0, "one": 1}

	n, ok := m["zero"]
	fmt.Println(n, ok)

	n, ok = m["notFound"]
	fmt.Println(n, ok)
}
