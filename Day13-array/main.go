// array class

package main

import "fmt"

func main() {
	// normal array declaration
	var arr [2]int

	// shorthand array declaration
	ar := [50]int{2, 3} // this is also okay: ar := [2]int{2, 3}

	fmt.Println(ar)
	fmt.Println(len(ar))
	fmt.Println(cap(ar))

	// arr = append(arr, 2)

	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
