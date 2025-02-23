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
