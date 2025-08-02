package main

import "fmt"

type User struct {
	id     int
	name   string
	age    int
	gender string
}

func main() {
	u1 := User{
		id:     1,
		name:   "mahmud",
		age:    25,
		gender: "male",
	}
	fmt.Println(u1)

	u1.age = 26

	fmt.Println(u1)
}
