package main

import "fmt"

func main() {
	m := map[string]int{"zero": 0, "one": 1}
	fmt.Println(m["zero"])
	fmt.Println(m["notFound"])
}
