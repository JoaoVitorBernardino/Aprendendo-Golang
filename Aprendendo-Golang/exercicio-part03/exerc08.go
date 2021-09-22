package main

import "fmt"

func main() {
	x := 3
	switch {
	case x == 0:
		fmt.Println("sono nível 0")
	case x == 1:
		fmt.Println("sono nível 1")
	case x == 2:
		fmt.Println("sono nível 2")
	case x == 3:
		fmt.Println("sono nível 3")
	}
}
