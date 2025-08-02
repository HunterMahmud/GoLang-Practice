package main

import "fmt"

/*
Function:

1. Parameter vs Argument
2. First order function
	i. Standard function or named function
	ii. Anonymous function
	iii. IIFE -> Immediately Invoke Function Expression
	iv. Function expression or function as a variable
3. Higher order function (receive a function or return a function or both)
	few names of higher order function:
		i. first class function
		ii. first class citizen
*/

// 1. standard or named function
func add(x, y int) int { // parameter receiving
	return x + y
}

// 5. higher order function -> this function can receive or
func higerFunction(x int, y int, sum func(x int, y int) int) int {
	var res int = sum(x, y)
	fmt.Println("Result inside higher order function: ", res)
	return res
}
func main() {
	fmt.Println(add(3, 4)) // argument passing

	//2. annonymous or IIFE -> Immediately Invoked Function Expression
	func() {
		fmt.Println("This will call imidiately.")
	}()

	// 4. functino expresssion or assigned function in a variable
	sum := func(x, y int) int {
		return x + y
	}

	var x int = sum(8, 9) // argument passing
	fmt.Println(x)

	var res int = higerFunction(4, 5, add)
	fmt.Println("Result outside higher order function: ", res)

}

// 3. init function
func init() {
	fmt.Println("This is init function and its called by computer. And it will execute first.")
}
