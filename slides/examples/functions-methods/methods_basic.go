package main

import "fmt"

type User struct{ Firstname, Lastname string }

// simple function
func Fullname(u User) string {
	return u.Firstname + " " + u.Lastname
}

// method on User
func (u User) Fullname() string {
	return u.Firstname + " " + u.Lastname
}

func main() {
	u1 := User{"Rob", "Pike"}

	// calling function
	fmt.Println(Fullname(u1))

	//calling method
	fmt.Println(u1.Fullname())

	var u2 User
	fmt.Println(u2.Fullname())
}
