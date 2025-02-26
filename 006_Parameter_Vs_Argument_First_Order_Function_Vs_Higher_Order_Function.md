## Parameter Vs Argument | First Order Function Vs Higher Order Function

#### Parameter Vs Argument

```go
 package main

import "fmt"

func add(a int, b int) { // Received value =>  parameter a and b
	c := a + b
	fmt.Println(c)
}
func main() {
	add(10, 20) // Pass value => Argument 10 and 20 
}

```

### First Order Function

- Standard function or named function 
- Anonymous function
- IIFE
- function expression

Note:

- First order function and Higher order function 
- come from functional programming paradime 
- example: haskel
- math -> logic (discreate mathematics)
- 1) First order logic
- 2) Higher order logic
- `Logic`
- 1) Object (people, animal, car etc)
- 2) Property (color (heair), student)
- 3) Relation 
- Rule: All customer must pay their pizza bills
- Tutul is student
- Apple is red
- Tutul is taller than Rakib.

####  First order logic 

- We can create any statement using logic rules
- All students must weare heir uniforms 

#### Higher order logic
 
 - Object properties, realation with rules that why it is heigher ordered 
 - complex logic apply 
 - Any rules that applies to all customers must also apply to tutul 
 - a rule: all the customers must pay tips to the waiters.

### Higher order function in GO

- 3 rules has to be applied 
- any one of the following 3 
- 1. Parameter -> function
- 2. function return
- 3. both 
  

#### function take as argument 
```go
package main

import "fmt"

// function as input parameter
func processOperation(a int, b int, op func(x int, y int)) {
	// op function execute here
	op(a, b)
}

// add and op similar as paramter list and type
func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	processOperation(10, 20, add)
}

```

#### function return from higher order function 

- function to function return
- when function return mention the return type
- anonymous function return

```go 
package main

import "fmt"

// function to function return
// when function return mention the return type
// anonymous function return
func call() func(x int, y int) {
	return add
}

// add and op similar as paramter list and type
func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	sum := call() // function asign called function expression

	sum(3, 4) // function call with argument
}


```

- collback function
- in higher order function we pass as paramter of function the function callback function
- higher order function a pass function called as callback function

--- 
- heigher order function or first class function
- if any function asing into valirable called as first class function
- first class citizen whaterver asing into variable 



