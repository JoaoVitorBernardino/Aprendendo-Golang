package main

import "fmt"

var x int = 10

//var x int
func main() {
	//x = 10
	//var x int = 10
	//x = 20.5 //erro pois é um valor float
	fmt.Println(x)
	fmt.Printf("x: %v, %T", x, x)
}
