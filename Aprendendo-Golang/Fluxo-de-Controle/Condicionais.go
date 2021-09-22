package main

import "fmt"

func main() {
	x := 10
	if x > 100 {
		fmt.Println("x é maior que cem")
	} else if x < 10 {
		fmt.Println("x é menor que dez")
	} else {
		fmt.Println("x não é menor que dez e nem maior que cem")
	}
}
