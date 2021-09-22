package main

import "fmt"

type teste int

var h teste = 10

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", h)
	//h = x
}
