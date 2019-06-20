package main

import "fmt"

type User struct{ Firstname, Lastname string }

func Fullname(u User) string {
	return u.Firstname + " " + u.Lastname
}

func (u User) Fullname() string {
	return u.Firstname + " " + u.Lastname
}

func main() {
	u1 := User{"Rob", "Pike"}
	fmt.Println(Fullname(u1), u1.Fullname())

	var u2 User
	fmt.Println(u2.Fullname())
}
