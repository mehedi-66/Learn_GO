package main

import "fmt"

func main() {
	// function expression or Assign function in valiable

	add(3, 4) // function expression not call before define
	// bud global function call any where
	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(2, 3)

}
