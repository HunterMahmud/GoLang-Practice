package main

import (
	"fmt"

	"abc/mathlib"
)

var p = 10

func main() {
	var (
		a = 10
		b = 20
	)

	sum := mathlib.Add(a, b)
	fmt.Println(sum)

	if true {
		var p = 20
		fmt.Println(p)
	}
	fmt.Println(p)
}
