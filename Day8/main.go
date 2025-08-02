package main

import (
	"fmt"

	"matlib.com/matlib"
)

func main() {
	var x int = add(5, 7)
	fmt.Println(x)

	fmt.Println(matlib.Add(2, 3))
}
