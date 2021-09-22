package main

import "fmt"

func main() {
	qualEsporte := "volei"
	switch qualEsporte {
	case "futebol":
		fmt.Println("meu esporte favorito é futebol")
	case "basquete":
		fmt.Println("meu esporte favorito é basquete")
	case "volei":
		fmt.Println("meu esporte favorito é volei")
	case "skate":
		fmt.Println("meu esporte favorito é skate")
	case "natação":
		fmt.Println("meu esporte favorito é natação")
	}
}
