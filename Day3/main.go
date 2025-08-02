package main

import "fmt"

var p = 100

const a = 10

// this is closure function
func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age: ", age)
	show := func() {
		money = money + p + a
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
	incr2()
}

func main() {
	call()
}
