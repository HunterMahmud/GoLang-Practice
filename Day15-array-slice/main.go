package main

import "fmt"

func main() {
	str := [6]string{"This", "is", "a", "Go", "interview", "questions."}

	sliceArr := str[1:5]

	fmt.Println(str)
	fmt.Println(sliceArr) // ptr:1, len:4, cap:5

	fmt.Println(len(sliceArr))
	fmt.Println(cap(sliceArr))

	arr := make([]int, 3, 5)

	fmt.Println("arr:", arr, "len:", len(arr), "cap:", cap(arr))
	arr = append(arr, 4)
	fmt.Println("arr:", arr, "len:", len(arr), "cap:", cap(arr))

	var s []int           // slice
	fmt.Println(s == nil) // true

	sl := []int{} // not slice

	fmt.Println(sl == nil) // false
}

/*
slice decalration:
	1. array to slice
	2. slice to slice
	3. slice literal (array declare but just size is missing) like: str := []string{"This", "is", "a", "Go", "interview", "questions."}
	4. using make function s := make([]int, 3) -> first one is type, second one is size
	5. using make function with capacity s := make([]int, 3, 6) first one is type, second one is size, third one is capacity
	6. empty slice or nil slice: var s []int

Golang slice underlying array capacity increase rules in append mode:
	1. Insufficient Capacity - Small Slices (current capacity < 1024): double or 2 times.
	2. Insufficient Capacity - Large Slices (current capacity >= 1024): 25% or 1.25 times.

*/
