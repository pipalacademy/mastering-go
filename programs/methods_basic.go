package main

import "fmt"

type user struct {
	firstname, lastname string
}

func fullname(u user) string {
	return u.firstname + " " + u.lastname
}

func (u user) fullname() string {
	return u.firstname + " " + u.lastname
}

func main() {
	rob := user{"Rob", "Pike"}
	fmt.Println(fullname(rob) == rob.fullname())
	fmt.Println(fullname(rob), rob.fullname())

}
