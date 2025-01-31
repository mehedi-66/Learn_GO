package main

import "fmt"

func add(x int, y int) int {
    return x + y;
}

func addAdnMul(x int, y int) (int, int) {

	return (x + y), (x*y);
}

func main(){
	a := 10;
	b := 20;

	x := add(a, b);
	fmt.Println( x);
	p, q := addAdnMul(a, b);
	fmt.Println( p, q);

}