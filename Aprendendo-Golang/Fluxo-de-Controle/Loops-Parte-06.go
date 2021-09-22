package main

import "fmt"

func main() {
	x := 10
	if x < 100 {
		fmt.Println("teste 1")
	}

	if !(x < 100) { // ! é uma negação
		fmt.Println("teste 2")
	}

	if y := 20; !(y < 100) {
		fmt.Println("teste 3")
	}
}
