package main

import "fmt"

type User struct{ Firstname, Lastname string }

func (u *User) Fullname() string {
	fn := (*u.Firstname) + (*u.Lastname)

	return fn
}

func main() {
	u := User{"Rob", "Pike"}
	up := &u
	fmt.Println(up.Fullname())
	fmt.Println(u.Fullname())
}
