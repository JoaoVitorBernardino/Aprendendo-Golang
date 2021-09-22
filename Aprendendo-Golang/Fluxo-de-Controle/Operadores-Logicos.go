package main

import "fmt"

func main() {
	x := 6
	if x == 2 || x == 3 {
		fmt.Println("X é 2 ou 3")
	} else if x%2 == 0 && x%3 == 0 {
		fmt.Println("é múltiplo de 2 e de 3")
	}
}
