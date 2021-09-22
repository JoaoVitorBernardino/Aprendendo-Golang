package main

import "fmt"

var x interface{}

func main() {
	/*x = -150
	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float")

	}*/
	switch x := 2; {
	case x == 1:
		fmt.Println("é 1")
	case x == 2:
		fmt.Println("é 2")
	case x == 3:
		fmt.Println("é 3")
	case x == 4:
		fmt.Println("é 4")

	}
}
