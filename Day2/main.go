package main

import "fmt"

func OperationalFunction(x int, y int, op func(a int, b int)) {
	op(x, y)
}

func add(a, b int) {
	fmt.Println(a + b)
}
func main() {
	OperationalFunction(2, 4, add)
}
