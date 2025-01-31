package main

import "fmt"

// Standard or name function
func add1(a int, b int) int {
	return a + b
}
func main() {

	fmt.Println("Hi I am mehedi")
	ans := add1(1, 2)
	fmt.Println(ans)
	// no name function called as anaomyous function
	// Imidately Invoced function Expression
	// IIFE
	func(a int, b int) {
		c := a + b
		fmt.Println("The sum of a and b is", c)
	}(1, 3)

}

func init() {
	fmt.Println("This is fun initialization first")

}
