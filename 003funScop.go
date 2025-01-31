package main

import "fmt"

func add(x int, y int) {
	var result = x + y
	fmt.Println("The sum is", result)
}
func main() {
	var a = 30
	var b = 20

	add(a, b)

}
