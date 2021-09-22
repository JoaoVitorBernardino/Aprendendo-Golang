package main

import "fmt"

func main() {
	x := 10
	if x > 10 {
		fmt.Println("X é maior que 10")
	} else if x == 10 {
		fmt.Println("X é igual a 10")
	} else {
		fmt.Println("X é menor que 10")
	}
}
