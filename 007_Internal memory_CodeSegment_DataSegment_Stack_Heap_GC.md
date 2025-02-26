
## Internal memory | Code Segment | Data Segment | Stack | Heap | GC

- Go Internal Memory 
- code Segment
- Data Segment
- Stack
- Heap

--- 

```go
package main

import "fmt"

var a = 10

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}
func main() {
	add(5, 4)
	add(a, 3)
}

func init() {
	fmt.Println("Hello")
}

```

---

- In the RAM when go program run that time 
- all segment memeory alocaiton can be increase depend on like dynamic memory 

- CodeSegment 
- DataSegment
- Stack 
- Heap 
- memeory alocated by go program
- Heap memeory managed by GC means Garbase collector 

--- 

- global memeory called variable all things  store into DataSegment
- codeSegment all function store 

--- 

- when program run 
- first `init()` function call then `Stack segment` alocated for exection
- second main() function call if not present through error 
    or stackSegment alocated for execution
- In stack alocated memeory called `stack-Frame`

--- 

## [Golang] 025 - End of Internal Memory

- if a function has inside  function or function expression where it store ?? 

- `Important`

```go
package main

import "fmt"

var a = 10
var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}
func main() {
	call()

	fmt.Println(a)
}

func init() {
	fmt.Println("Hello")
}

```

- Phase 2 in go program run
- 1. compilation phase
- 2. execution phase

- go run main.go => compile it => main.exe => ./main
- go build main.go => compile it => main.exe
--- 
### comiler Phase
1. Complier when build go program that time `Binary` file generated 
- 00000010101
01000100101
0101010

        CodeSegment `const` store 
        // which are read only that are goes into code section

        a = 10
        call = func {...}
        add = func {...}
        main = func {...}
        init = func {...}

-
        dataSegement `p` store

- 
### execution phase

- when exectution start created into RAM
- codeSegment 
- dataSegment
- Stack
- Heap 

- when execution start function alocated `stack frame`
- `function add codeSegment before alocation into codeSegment` 
- `add function reference` to stack frame when alocation
- inside call function only access the add function only 
- when we calling add(2,3) function that time stack frame is created 






