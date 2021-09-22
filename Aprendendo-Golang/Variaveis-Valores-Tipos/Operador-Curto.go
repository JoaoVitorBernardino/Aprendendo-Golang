package main

import "fmt"

var y = "Olá, bom dia"

func main() {
	x := 10

	//fmt.Println(x, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x: %v, %T\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x: %v, %T\n", z, z)
}
