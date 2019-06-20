package main

import "fmt"

type User struct{ Firstname, Lastname string }

// func (u User) Fullname() string {
// 	return u.Firstname + " " + u.Lastname
// }

func (u *User) Fullname() string {
	// return (*u).Firstname + " " + (*u).Lastname
	return u.Firstname + " " + u.Lastname
}

func main() {
	u := User{"foo", "bar"}
	up := &u
	fmt.Println(up.Fullname())

}
