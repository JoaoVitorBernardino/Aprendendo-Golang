package main

import "fmt"

func main() {
	x := `"oi bom dia\n

			como vai?\t    espero \"que\" 	
											bem."`
	y := "bom dia"
	z := fmt.Sprint(x, " ", y)
	fmt.Println(x)
	//fmt.Print(x)
	//fmt.Printf(x)
	fmt.Println(y)
	fmt.Println(z)
}
