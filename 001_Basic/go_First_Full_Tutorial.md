### Six Main Points about GO

<img src="img/001.png">

#### Two things => Modules, Packages

- Collection of files under Folder calld Package
- Collection of Package called Module 

<img src="img/002.png">

- Create Module 
- => go mod init NameModule
- go.mod file is create with version number

<img src="img/003.png">

- Every go file uder the package 
- Create binary file of Compiled program 
- => go build FileName.go
- => go run FileName.go  // build and run 

<img src="img/004.png">

- int16 int32 int64 this are how many bit and number we can store 
- If we store maximum number then overflow 
- gives wrong result (runtimes )

<img src="img/005.png">

- uint16 uint32 uint64 only sotee positive integer 

        Example:
            int8: (-128, 127)
            uint8: (0, 255)

<img src="img/006.png">

- Arithmetic operations gives wrong diffrent data type 

<img src="img/007.png">

        type cust one of the number

        var resut float32 = floatNum32 + float32(intNum32)

- ` To Direct formationg `` ` 

<img src="img/008.png">

- concatinationg string by + 
- "Hello" + "World"
- `len("test")` => 4 

<img src="img/009.png">

<img src="img/010.png">

- Bolenan true and flase 

         var myBool bool = true
         fmt.Println(myBool)  >>> true


- if we not initialize variable then asing default value asing by compiler 

<img src="img/011.png">

<img src="img/012.png">

- empty string
<img src="img/013.png">

- bool it is false


#### Short hand or dynamicaly asing 

- var myVar = "test"  => string type compiler asing 
- myVar := "test"
- Fancy way 
- var var1, var2 = 3, 4 
- var1, var2 := 3, 4
- whem function return that time variable type declaration good practice

<img src="img/014.png">

- constant is value does not change 
- constant when declaration must initialize the value 
<img src="img/015.png">

### Funxtion and Control Structure 

<img src="img/016.png">

<img src="img/017.png">

- integerDivision

<img src="img/018.png">

- return multiple results (int, int)

<img src="img/019.png">

<img src="img/020.png">


#### Return error from function 

- error type variable default value is `nil`
- return error from the function 
- and grab the erro which is return 

<img src="img/022.png"> 
<img src="img/021.png"> 

- multiple conditions 
<img src="img/023.png"> 

- In go && operation first false rest skiped 
<img src="img/024.png"> 

- In go || operation true rest skiped
<img src="img/025.png"> 

- switch statements 

<img src="img/026.png"> 

<img src="img/027.png"> 

### Array, Slices, Maps and Loop


<img src="img/028.png"> 

#### Array 

<img src="img/029.png"> 

- intArr[1:3] => give index 1, 2 value less one index of 3
- change value intArr[0] = 340

<img src="img/030.png"> 

- 3 size array gives 12 bytes of contiguous memory location 
- print memory location 

<img src="img/031.png"> 

- initialize array by this 
<img src="img/032.png"> 

- initialize array by shorthand  
<img src="img/033.png"> 

- Compiler gives the size in case here [3]
<img src="img/034.png"> 


#### Slices

<img src="img/035.png"> 
<img src="img/036.png"> 

- capacity increase by 2 when no room to append value to slice array 

<img src="img/037.png"> 

- when length of the elment index acess then give outoffrange error 

<img src="img/038.png"> 

- Multiple elements insert at a time  

<img src="img/039.png"> 

- slice create by make function 
- capacity is optional 8

<img src="img/040.png"> 


### Map 

- map is key value pair data structure

<img src="img/041.png"> 

- key type is string and vlaue type is unsigned int8

- if key is not present then default value is of that time give in our case 0

- for delete value from map has 
- => delete(myMap, "mehedi"
)
<img src="img/042.png"> 

        var age, ok = myMap["mehedi"]

        // return ok variable map return true or false depending key present or not 


<img src="img/043.png"> 

- iterate throught is map 
- range keyword array, slice, map iterate 

<img src="img/044.png"> 

- i is index and v is value here 

<img src="img/045.png"> 

- while loop in go

<img src="img/046.png"> 

- infinity loop with break key word

<img src="img/047.png"> 

- for loop in go

<img src="img/048.png"> 

- a progream how many time to n number of elements insert into slice\


<img src="img/049.png"> 

- Total time without Preallocation: 15.6682ms
- Total time with Preallocation: 2.4033ms


### String, Runes, Bytes 

<img src="img/050.png"> 

- non ASCII character in string 
- looping skip 2 
- UTF8 in string  "string" by 8 bit binary encoding
- ASCII in string by 7 bit binary encoding

<img src="img/051.png"> 
<img src="img/052.png"> 
<img src="img/053.png"> 
<img src="img/054.png">

-  "" to give `[]rune` 

<img src="img/055.png">


                var myRune = 'a'
                fmt.Printf("\nmyRune = %v", myRune) // 97


- Every time new string is created 
- string is imutable that that means by index `myString[2]` value change 
- this looping way create string is ineficient 

<img src="img/056.png">

- using string builder create string more faster 
- and efficient 

<img src="img/057.png">

### Struct and Interface in GO

<img src="img/058.png">

- user define type 
- here type is struct 
- struct contain mixure type of vairables

<img src="img/059.png">

- intialize and  print 

<img src="img/060.png">

- omit the fild name has to order flowing
- or name.fildName access 

<img src="img/061.png">

- struct inside struct example
- struct variale use inside of struct example 

<img src="img/062.png">

- direct struct using inisde struct 

<img src="img/063.png">

- Anonymous struct and inisized when it create 
- it is not reuseable 

<img src="img/064.png">

- Struct has called method that can use struct
- here method like class mehtod like other programming language concept 
- thes are method that is tried to struct 
- access the the struct variable 
- method just like fuction 
- `like class (struct) instance create and called is method` 

<img src="img/065.png">

```go
package main
import "fmt"

type myStruct struct{
        Name string
	Age int
}
func (e myStruct) printName(){
        fmt.Println("Hi "+ e.Name)
}
func main() {
	var mySt myStruct = myStruct{Name: "Mehedi"};
	mySt.printName()
	fmt.Println(mySt.Name)
}
```

- Any Engine Miliage calculation 
- that time iterface is comming 

<img src="img/067.png">

- Interface 
- using this take any parameter 
- take any type of struct as paraper 
- and dynamicall call the function for calucalete distance 

<img src="img/068.png">

### Pointer in Go

<img src="img/069.png">

- pointer store address of memory location 

<img src="img/070.png">

- when `var p *int32` nil store into memeory address where p memeory alocated 
- that means null store if not assing anything s
-  

<img src="img/071.png">

- when `*p` vlaue is assing the p address refereance location value is updated 

<img src="img/072.png">

- `&i` to store i variable memeory address asing to a pointer 
- if there are refereancing to each other 
- one's value is changed then both see the value is updated 
- 

<img src="img/073.png">
<img src="img/074.png">

- if normal valriable not copy not share refereance it is value type 
- on the other hand slices used the refereance under the hood 
- if we copy then both array value is changed acordingly 


<img src="img/075.png">

#### pointer with function 

- if we pass array into function then value type is passed to function array parameter
- and both array address prointer is differet on thing1 and thing2

<img src="img/076.png">

<img src="img/077.png">

- Function pass arguments `&` and paramter `*`
- pass as refereance same memory location share 
- copay value into new array takes time we can reduce time
- if any vlaue change that going to refalact both arrays

<img src="img/078.png">


### concurancy and Gorutine in Go 

<img src="img/079.png">

- lunch multiple functions
- concurancy working that function 
- concurancy != Parallel execution 

#### concurancy

- single CPU execution task by switching context 
- non blocking execution task
- when database call by 3 second 
- that time also task2 execute by context switching


<img src="img/080.png">

#### Parallel execution

- another way get concurancy by parallel execution
- Two task and dedicated two CPU `Core`
- symmetricaly execution task1 and task2 by two dedicated CPU core
- go achive some level of parallel execution by `goroutine`

<img src="img/081.png">

- DB call simulte without go routines
- the time it taken is 5.03 s


<img src="img/082.png">

- lets database call concurently by `go`
- when `go dbCall(i)` the program finish without any output 
- why the task is execute in background 
- we have to wait for the task to finish
- wating group come into the picture `sync` libray can help in this case 
- `sync.WaitGroup{}` like a counter 

- wg.Add(1) set counter of 
- wg.Done() decrement the counter
- wg.Wait() all the counter are back to zero
- that means all the task are execute 

- Total Time of execution: 1.2 s significantly improve performance


<img src="img/083.png">

- some unexpected result store into our slice 
- because of a time differenct processes or thread  access the same memory curept the memoery location the slice to store the result

<img src="img/084.png">


- using `sync.Mutex{}` macanizome the lock and ulock 
- access the value and update the value using thread can safely be
- if one theread is lock a varaible the other threead has to be wait until that is unlocked 

<img src="img/085.png">

- full lock and read lock 

<img src="img/086.png">

- Task function not have that much work then less time 
- if some amount of work is required few times
- still so fast 
- example 

<img src="img/087.png">

- Task is faster depending on Number of CPU Core (8)

<img src="img/088.png">

### Channels in Go-Routine

<img src="img/089.png">

- channel in go routine how the infomation pass between go routines 
- Hold data
- Thread safe
- Listen for data 
- when listen holding code execution 

<img src="img/090.png">

- using `make(chan int)` or string any type of data can create channel for pass
- channel like hold array or buffere data 
- `<-` store and retrive data from chanenl

<img src="img/091.png">

```go 
package main

import "fmt"

func main() {
        var c = make(chan int)
	c <- 1  // here asing value code here wait for read data from it 
	         // give deadlock error shows here 
	var i = <-c

	fmt.Println(i) // 1
}

```

- `Proper way to write channel with go routine`

- first create channel
- then pass to go routine function call bacakgroud function execute
- then immediately fmt println(<-c) 
here wait channel data recive by which go routine function send data
- then print properly 

```go
package main

import "fmt"

func main() {
	var c = make(chan int)
	go process(c) // bacakgroud function execute
	fmt.Println(<-c) // send value print and wait for data send data from go routine function execute
}

func process(c chan int) {
	c <- 123 // assign send data 
}

```

- after printing value from channel
- deadlock again happened 
- after sending all the values 
- channel process need to `close()`

 <img src="img/093.png">

 ```go 
package main

import "fmt"

func main() {
        var c = make(chan int)
	go process(c)

	for i := range c {

		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	//close(c)
}

 ```

 - Buffer channel means [1,2,4] like store value into go routine fucsiont 
- into buffer array 
- then main function slowly read from the buffer array 

<img src="img/094.png">

- `Select` to check which chanecl gives value 
- sales checking price and Max_Chicken_Price threshold

<img src="img/096.png">

- using two channels which chanel send message write that 
- select like a if satement 

<img src="img/095.png">

### Generics in Go

<img src="img/097.png">

- for same task using function different function we have to write 
- if we pass a generic type variable worked done 

<img src="img/098.png">

- insted of passing value 
- we are going to send type `[T int | string]`

<img src="img/099.png">

- we can use `any` type 

<img src="img/100.png">

- load Json and struct example 
- and also generic 
- unmarsall 

<img src="img/101.png">

<img src="img/102.png">

- struct and generic example

<img src="img/103.png">


