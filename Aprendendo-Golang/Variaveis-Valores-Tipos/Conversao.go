package main

import "fmt"

type tipo int

var t tipo = 10

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	//fmt.Printf("%T", t)

	x = int(t)
	fmt.Printf("%T\n", x)
}
