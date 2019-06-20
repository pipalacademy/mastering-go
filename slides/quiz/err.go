package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, _ := ioutil.ReadFile("some_non_existent.go")
	fmt.Println(content)
}
