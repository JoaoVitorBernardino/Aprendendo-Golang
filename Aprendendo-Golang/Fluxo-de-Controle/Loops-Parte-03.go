package main

import "fmt"

func main() {
	x := 0
	for ; x < 10; x++ {
		fmt.Println(x)
	}
	/*
		Loop infinito:
		for {
			fmt.println(x)
		}

		Loop condicional:
		for {
			if x < 10 {
				x++
				fmt.println(x)
			} else {
				fmt.println("saí do loop")
				break
			}
		}
		fmt.println("O loop está breakado!")
	*/
}
