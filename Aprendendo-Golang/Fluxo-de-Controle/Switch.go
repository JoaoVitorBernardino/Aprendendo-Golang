package main

import "fmt"

func main() {
	/*x := 5
	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")

	}*/
	quemTaNoEscritorioHoje := "joana"
	switch quemTaNoEscritorioHoje {
	case "zezinho", "maria":
		fmt.Println("Hoje quem ta no escritório é o time 1")
		//fallthrough // roda o caso de baixo também
	//case "marquinhos":
	//fmt.Println("Hoje quem ta no escritório é marquinhos")
	case "joana", "pedrinho":
		fmt.Println("Hoje quem ta no escritório é o time 2")
		//fallthrough
	//case "maria":
	//fmt.Println("Hoje quem ta no escritório é maria")
	default:
		fmt.Println("O escritório está vazio.")
	}
}
