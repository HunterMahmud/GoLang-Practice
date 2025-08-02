package main

import "fmt"

func print(numbers *[3]int) {

	fmt.Println(*numbers)
}

func main() {
	// arr := [6]int{1, 2, 3, 4, 5, 6}
	// fmt.Println("arr:", arr, "len:", len(arr), "capacity:", cap(arr))

	// // literal slice-> directly initialize
	// s3 := []int{2, 3, 4}
	// fmt.Println("slice3:", s3, "len:", len(s3), "capacity:", cap(s3))

	// // print(&arr)

	// // normal slice
	// s1 := arr[1:4] // 2,3,4 -> len=3 -> cap=5
	// // slice from slice
	// s2 := s1[1:2] // 3 -> len=1 -> cap=4

	// // using make function
	// s4 := make([]int, 3)
	// // using make function with capacity
	// s5 := make([]int, 3, 5)

	// // nil slice

	// s6 := []int{}

	// s6 = append(s6, 2)

	// fmt.Println("slice1:", s1, "len:", len(s1), "capacity:", cap(s1))
	// fmt.Println("slice2:", s2, "len:", len(s2), "capacity:", cap(s2))
	// fmt.Println("slice4:", s4, "len:", len(s4), "capacity:", cap(s4))
	// fmt.Println("slice5:", s5, "len:", len(s5), "capacity:", cap(s5))
	// fmt.Println("slice6:", s6, "len:", len(s6), "capacity:", cap(s6))

	// another code:
	// x := []int{}

	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))
	// x = append(x, 1)
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))
	// x = append(x, 2)
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))
	// x = append(x, 3)
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))
	// x = append(x, 4)
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))
	// x = append(x, 5)
	// x = append(x, 6)
	// x = append(x, 7)
	// x = append(x, 8)
	// // x = append(x, 9)

	// x = append(x, 1) // [1], len:1, cap:1
	// x = append(x, 2) // [1,2], len:2, cap:2
	// x = append(x, 3) // [1,2,3], len:3, cap:4
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))

	// y := x // [1,2,3], len:3, cap:4

	// x = append(x, 4) // [1,2,3,4], len:4, cap:4
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))
	// y = append(y, 5) // [1,2,3,5], len:4, cap:4
	// fmt.Println("slice-y:", y, "len:", len(y), "capacity:", cap(y))
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))

	// x[0] = 10 // [10,]
	// fmt.Println("slice-x:", x, "len:", len(x), "capacity:", cap(x))
	// fmt.Println("slice-y:", y, "len:", len(y), "capacity:", cap(y))

	// another code:
	x := []int{1, 2, 3, 4, 5} // [1,2,3,4,5], len:5, cap:5
	x = append(x, 6)          // [1,2,3,4,5,6], len:6, cap:10
	x = append(x, 7)          // [1,2,3,4,5,6,7], len:7, cap:10
	a := x[4:]                // [5,6,7], len:3, cap:6
	fmt.Println("the a is:", a, "len:", len(a), "cap:", cap(a))

	y := changeSlice(a) // [10,6,7,11], len:4, cap:6

	// fmt.Println("x", x, "len:", len(x), "cap:", cap(x)) // [1,2,3,4,10,6,7], len:7, cap:10
	// fmt.Println("y", y, "len:", len(y), "cap:", cap(y)) // [10,6,7,11], len:4, cap:6
	// fmt.Println("a", a, "len:", len(a), "cap:", cap(a)) // [10,6,7], len:3, cap:6

	fmt.Println(y[0:cap(y)])

	// p := make([]int, 3, 6)

	// fmt.Println(p[0:6])

}

func changeSlice(a []int) []int {
	a[0] = 10         // [10,6,7], len:3, cap:6
	a = append(a, 11) // [10,6,7,11], len:4, cap:6
	return a
}
