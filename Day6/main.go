package main

import "fmt"

func calculation() (result int) {
	fmt.Println("first", result)
	show := func() {
		result += 10
		fmt.Println("defer:", result)
	}
	defer show()

	result = 5
	fmt.Println("second:", result)
	return
}
func calc() int {
	result := 0
	fmt.Println("first", result)
	defer func() {
		fmt.Println("defer anon: ", result)
	}()
	show := func() {
		result += 10
		fmt.Println("defer:", result)
	}
	defer show()

	result = 5
	fmt.Println("second:", result)
	return result
}
func main() {

	a := calculation()

	fmt.Println("calculation: ", a)

	b := calc()
	fmt.Println("calc: ", b)
}
