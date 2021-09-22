package main

import "fmt"

type teste int

var h teste

func main() {
	fmt.Println(h)
	fmt.Printf("%T\n", h)
	h = 42
	fmt.Println(h)
}
