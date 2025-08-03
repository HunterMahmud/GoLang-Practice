package main

import (
	"fmt"
	"math"
)

func print(numbers ...[]int) { // vairadic funciton  numbers ...int if we want only numbers

	numbers[0] = append(numbers[0], 12, 23, 34)
	fmt.Println(numbers[0])
	fmt.Println(len(numbers[0]))
	fmt.Println(cap(numbers[0]))
	fmt.Println("--------------")
	fmt.Println(numbers[1])
	fmt.Println(len(numbers[1]))
	fmt.Println(cap(numbers[1]))
}

func main() {
	// var x = []int{1,2,3} // or x := []int{1,2,3} slice
	// x := [1024]int{}
	y := []int{1, 2}
	z := []int{3, 4, 5}
	print(y, z)
	fmt.Println("Maximum uint64 value:", uint64(math.MaxUint64))
	fmt.Println(float32(math.MaxFloat32))
	fmt.Println(float64(math.MaxFloat64))
}
