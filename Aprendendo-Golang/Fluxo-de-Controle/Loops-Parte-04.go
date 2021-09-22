package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			//faz isso quando o número não é par
			continue //quebra a interação específica, diferente do break que quebra o loop por inteiro
		}
		fmt.Println(i)
	}
}
