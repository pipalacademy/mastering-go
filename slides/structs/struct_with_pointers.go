package main

import "fmt"

type User struct{ Firstname, Lastname string }

func (u *User) Fullname() string {
	// no need of accessing the fields like `(*u).field`. simply do u.Field
	// return (*u).Firstname + " " + (*u).Lastname

	return u.Firstname + " " + u.Lastname
}

func (u User) setFirstName1(fn string) {
	u.Firstname = fn
}

func (u *User) setFirstName2(fn string) {
	u.Firstname = fn
}

func main() {
	u := User{"Foo", "Bar"}
	up := &u
	fmt.Println(up.Fullname())
	fmt.Println(u.Fullname())

	u.setFirstName1("changeMe")
	fmt.Println(u) // Foo Bar

	u.setFirstName2("changeMe")
	fmt.Println(u) // changeMe Bar

}
