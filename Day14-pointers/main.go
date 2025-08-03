// pointers
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printNumbers(numbers [3]int) { // this will copy all the values of arr to numbers
	fmt.Println(numbers)
}

func printNumbersWithPointers(numbers *[3]int) { // this will only copy the first addres of the arr
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	x := 20

	fmt.Println("value of x:", x)

	p := &x // p is a pointer that it pointing the address of x

	fmt.Println("address of x:", p)
	fmt.Println("value of x using its pointer:", *p)

	// array
	arr := [3]int{1, 2, 3}

	printNumbers(arr)
	printNumbersWithPointers(&arr)

	mahmud := User{
		Name: "mahmud",
		Age:  26,
	}

	ptr := &mahmud

	fmt.Println(*&*&ptr.Age) // moja kore print korci.. just ghurai firai aki jinis: ptr.Age => *&ptr.AGe => *&*&ptr.Age

}
