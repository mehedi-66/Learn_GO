package main

import "fmt"

/* Declare a single variable */
var a int

var (
	b bool
	c float32
	d string
)

func main() {

	a = 42            // Assign single value
	b, c = true, 32.0 // Assign multiple valeus
	d = "string"      // sring must contain double quotes
	fmt.Println(a, b, c, d)

	// Initialize and assign to a single variable

	e := 42            // Initialize and assign to a single variable
	f, g := true, 32.0 // Initialize and assign to multiple variables
	h := "string"
	fmt.Println(e, f, g, h) // 42 true 32 string

}
