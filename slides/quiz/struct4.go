package main

import "fmt"

type Employee struct {
	Name    string
	Team    string
	Manager Employee
}

func main() {
	e := Employee{}
	fmt.Println(e)
}
