package main

import "fmt"

// variable shadowing
func main() {
	var a int = 10
	func() {
		var a int = 20
		fmt.Println(a)
		func() {
			var a int = 30
			fmt.Println(a)
		}()
	}()
	fmt.Println(a)
}
