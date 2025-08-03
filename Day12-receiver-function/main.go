package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

// receiver function - only works for struct
func (usr User) printDetails(x int) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
	fmt.Println("but he is really:", x)
}

func main() {
	u1 := User{
		Name: "Mahmud",
		Age:  26,
	}

	printUserDetails(u1)

	u1.printDetails(25)

}
