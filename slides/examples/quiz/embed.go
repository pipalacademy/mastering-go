package main

import "fmt"

type Team struct {
	Name string
	Size int
}

type Employee struct {
	Name string
	Team
}

func main() {
	e := Employee{"foo", Team{"blues", 10}}
	fmt.Println(e.Name)

	e1 := Employee{}
	e1.Team = Team{"blues", 10}
	fmt.Println(e1.Name)
}
