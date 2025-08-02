package main

import "fmt"

type User struct {
	name string
	age  int
}

func printUserDetails(usr User) {
	fmt.Println("user name:", usr.name)
	fmt.Println("user age:", usr.age)
}

// receiver function
func (usr User) printDetails(a int) {
	fmt.Println("user name:", usr.name)
	fmt.Println("user age:", usr.age)
	fmt.Println("given value", a)
}

func main() {
	user1 := User{
		name: "mahmud",
		age:  25,
	}

	printUserDetails(user1)
	user2 := User{
		name: "Abdullah",
		age:  23,
	}
	user2.printDetails(10)
}
