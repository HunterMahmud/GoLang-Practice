package main

import (
	"fmt"
	"mahmud/addition"
)

func call() int { // unnamed return value or function
	i := 0

	fmt.Println("first: ", i)

	defer fmt.Println("second:", i)

	i++
	fmt.Println("third:", i)
	defer fmt.Println("Fourth:", i)

	return i
}

func sum(x int, y int) (result int) { // named return value or function
	result = x + y // don't have to write like result := x + y, already have result
	return         // we can also use return result or just keep only return
}

func calculation() (result int) {

	fmt.Println("first: ", result)

	show := func() {
		result += 10
		fmt.Println("defer: ", result)
	}

	defer show()
	result = 5

	return // will return 15

}

func calc() int {
	result := 0

	fmt.Println("first: ", result)

	show := func() {
		result += 10
		fmt.Println("defer: ", result)
	}

	defer show()
	result = 5

	return result // will return 5 without named it snapshot the return value. So if it changes after that it will not effect to return.
}

func main() {
	fmt.Println("call:", call())
	fmt.Println("sum:", sum(2, 3))
	fmt.Println("with named return func: ", calculation())
	fmt.Println("without named return func: ", calc())
	fmt.Println("another file accessing: ", Add(2, 3))
	fmt.Println("another folder accessing: ", addition.Addition(4, 3))

}

/*
defer functin follows the stack's behavior: FILO -> First in last out.

*/
